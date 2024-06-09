package semantic

import (
	"fmt"

	"github.com/authzed/spicedb/pkg/namespace"
	core "github.com/authzed/spicedb/pkg/proto/core/v1"
	"github.com/authzed/spicedb/pkg/schemadsl/compiler"
)

type Relation struct {
	name       string
	visibility Visibility
	inType     *Type
	extensions []*ExtensionReference
	body       RelationExpression
}

func NewRelation(name string, visibility Visibility, body RelationExpression) (*Relation, error) {
	r := &Relation{name: name, visibility: visibility, body: body}

	return r, nil
}

func (r *Relation) SpiceDBName() string {
	return fmt.Sprintf("p_%s", r.name)
}

func (r *Relation) VisibleTo(t *Type) bool {
	return true
}

func (r *Relation) ReferencedTypes() ([]*Type, error) {
	return r.body.ReferencedTypes(r)
}

func (r *Relation) ToZanzibar() ([]*core.Relation, error) {
	var relations []*core.Relation
	body, err := r.body.ToZanzibar(r)
	if err != nil {
		return []*core.Relation{}, err
	}
	relation, err := namespace.Relation(r.name, namespace.Union(body))
	if err != nil {
		return []*core.Relation{}, err
	}
	relations = append(relations, relation)

	relationTypes, err := r.ReferencedTypes()
	if err != nil {
		return relations, err
	}

	if len(relationTypes) > 0 {
		types := []*core.AllowedRelation{}
		for _, t := range relationTypes {
			types = append(types, namespace.AllowedRelation(t.SpiceDBName(), compiler.Ellipsis))
		}
		relations = append(relations, namespace.MustRelation(r.SpiceDBName(), nil, types...))
	}

	return relations, nil
}

type RelationExpression interface {
	ToZanzibar(*Relation) (*core.SetOperation_Child, error)
	ReferencedTypes(*Relation) ([]*Type, error)
}

type SelfRelationExpression struct {
	referencedType *TypeReference
	cardinality    Cardinality
}

func NewSelfRelationExpression(referencedType *TypeReference, cardinality Cardinality) *SelfRelationExpression {
	return &SelfRelationExpression{
		referencedType: referencedType,
		cardinality:    cardinality,
	}
}

func (e *SelfRelationExpression) ToZanzibar(r *Relation) (*core.SetOperation_Child, error) {
	if !e.referencedType.IsResolved() {
		err := e.referencedType.Resolve(r.inType)
		if err != nil {
			return nil, err
		}
	}

	wrapper := namespace.ComputedUserset(r.SpiceDBName()) //SpiceDB-specific

	return wrapper, nil
}

func (e *SelfRelationExpression) ReferencedTypes(r *Relation) ([]*Type, error) {
	if !e.referencedType.IsResolved() {
		err := e.referencedType.Resolve(r.inType)
		if err != nil {
			return nil, err
		}
	}
	return []*Type{e.referencedType.instance}, nil
}

type ReferenceRelationExpression struct {
	relation    string  //Relation name on the same type
	subrelation *string //Optional relation on related objects
	//Note: when using subrelations, we need to validate visiblity AND it's tricky because the actual types may be many
	//Will need to recursively evaluate the body of `relation` for this type to find all SelfRelationExpressions
	// ..and do the same for any ReferenceRelationExpressions encountered to get the set of all types.
	//Then we can check if the named subrelation is accessible on all possible related types
	//..May skip for now.
}

func NewReferenceRelationExpression(relation string, subrelation *string) *ReferenceRelationExpression {
	return &ReferenceRelationExpression{
		relation:    relation,
		subrelation: subrelation,
	}
}

func (e *ReferenceRelationExpression) ToZanzibar(r *Relation) (*core.SetOperation_Child, error) {
	if e.subrelation == nil {
		return namespace.ComputedUserset(e.relation), nil
	}

	//Oh jeez. SpiceDB doesn't allow a permission on the left side of an arrow. Soooooooo, this doesn't work.
	//Instead! Let's:
	// Get a pointer to the relation
	// Add a method to the relation that gets all the descendent selfrelationexpressions (specifically, the relations it added)
	// For each of those such that they have the given relation and it's accessible, union them together
	relation, ok := r.inType.relations[e.relation]
	if !ok {
		return nil, ErrSymbolNotFound
	}

	relationTypes, err := relation.ReferencedTypes()
	if err != nil {
		return nil, err
	}

	for _, t := range relationTypes {
		subrelation, ok := t.relations[*e.subrelation]
		if !ok {
			continue
		}

		if !subrelation.VisibleTo(r.inType) {
			return nil, ErrSymbolNotAccessible
		}
	}
	return namespace.TupleToUserset(relation.SpiceDBName(), *e.subrelation), nil
}

func (e *ReferenceRelationExpression) ReferencedTypes(r *Relation) ([]*Type, error) {
	return []*Type{}, nil
}

type SetRelationExpression struct {
	kind  string
	left  RelationExpression
	right RelationExpression
}

func NewSetRelationExpression(kind string, left RelationExpression, right RelationExpression) *SetRelationExpression {
	return &SetRelationExpression{
		kind:  kind,
		left:  left,
		right: right,
	}
}

func (e *SetRelationExpression) ToZanzibar(r *Relation) (*core.SetOperation_Child, error) {
	leftbody, err := e.left.ToZanzibar(r)
	if err != nil {
		return nil, err
	}

	rightbody, err := e.right.ToZanzibar(r)
	if err != nil {
		return nil, err
	}

	var body *core.UsersetRewrite

	switch e.kind {
	case "union":
		body = namespace.Union(leftbody, rightbody)
	case "intersect":
		body = namespace.Intersection(leftbody, rightbody)
	case "except":
		body = namespace.Exclusion(leftbody, rightbody)
	}

	return namespace.Rewrite(body), nil
}

func (e *SetRelationExpression) ReferencedTypes(r *Relation) ([]*Type, error) {
	leftTypes, err := e.left.ReferencedTypes(r)
	if err != nil {
		return []*Type{}, err
	}

	rightTypes, err := e.right.ReferencedTypes(r)
	if err != nil {
		return []*Type{}, err
	}

	return append(leftTypes, rightTypes...), nil
}
