package semantic

import (
	"fmt"

	"github.com/authzed/spicedb/pkg/namespace"
	core "github.com/authzed/spicedb/pkg/proto/core/v1"
	"github.com/authzed/spicedb/pkg/schemadsl/compiler"
)

type Relation struct {
	name          string
	visibility    Visibility
	inType        *Type
	extensions    []*ExtensionReference
	body          RelationExpression
	generatedFrom *DynamicRelation
}

func NewRelation(name string, t *Type, visibility Visibility, body RelationExpression, generatedFrom *DynamicRelation) (*Relation, error) {
	r := &Relation{name: name, inType: t, visibility: visibility, body: body, generatedFrom: generatedFrom}

	return r, nil
}

func (r *Relation) SpiceDBName() string {
	return fmt.Sprintf("t_%s", r.name)
}

func (r *Relation) VisibleTo(t *Type) bool {
	return true
}

func (r *Relation) AddExtension(e *ExtensionReference) {
	e.namespace = r.inType.namespace
	e.onType = r.inType
	e.relation = r

	r.extensions = append(r.extensions, e)
}

func (r *Relation) ApplyExtensions() error {
	for _, e := range r.extensions {
		err := e.Apply()
		if err != nil {
			return err
		}
	}
	return nil
}

func (r *Relation) DirectTypeReferences() ([]*TypeReference, error) {
	return r.body.DirectTypeReferences(r)
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

	typeReferences, err := r.DirectTypeReferences()
	if err != nil {
		return relations, err
	}

	if len(typeReferences) > 0 {
		allowedRelations := []*core.AllowedRelation{}
		for _, ref := range typeReferences {
			if !ref.IsResolved() {
				err := ref.Resolve(r.inType)
				if err != nil {
					return relations, err
				}
			}
			if ref.all {
				allowedRelations = append(allowedRelations, namespace.AllowedPublicNamespace(ref.instance.SpiceDBName()))
			} else {
				var subrelation string
				if ref.subRelation != "" {
					subrelation = ref.subRelation
				} else {
					subrelation = compiler.Ellipsis
				}
				allowedRelations = append(allowedRelations, namespace.AllowedRelation(ref.instance.SpiceDBName(), subrelation))
			}
		}
		relations = append(relations, namespace.MustRelation(r.SpiceDBName(), nil, allowedRelations...))
	}

	return relations, nil
}

type RelationExpression interface {
	ToZanzibar(*Relation) (*core.SetOperation_Child, error)
	DirectTypeReferences(*Relation) ([]*TypeReference, error)
}

type SelfRelationExpression struct {
	referencedTypes []*TypeReference
	cardinality     Cardinality
}

func NewSelfRelationExpression(referencedTypes []*TypeReference, cardinality Cardinality) *SelfRelationExpression {
	return &SelfRelationExpression{
		referencedTypes: referencedTypes,
		cardinality:     cardinality,
	}
}

func (e *SelfRelationExpression) ToZanzibar(r *Relation) (*core.SetOperation_Child, error) {
	for _, t := range e.referencedTypes {
		if !t.IsResolved() {
			err := t.Resolve(r.inType)
			if err != nil {
				return nil, err
			}
		}
	}

	wrapper := namespace.ComputedUserset(r.SpiceDBName()) //SpiceDB-specific

	return wrapper, nil
}

func (e *SelfRelationExpression) DirectTypeReferences(r *Relation) ([]*TypeReference, error) {
	return e.referencedTypes, nil
}

type ReferenceRelationExpression struct {
	relation    string
	subrelation *string
}

func NewReferenceRelationExpression(relation string, subrelation *string) *ReferenceRelationExpression {
	return &ReferenceRelationExpression{
		relation:    relation,
		subrelation: subrelation,
	}
}

func (e *ReferenceRelationExpression) ToZanzibar(r *Relation) (*core.SetOperation_Child, error) {
	relation, ok := r.inType.relations.Get(e.relation)
	if !ok {
		return nil, fmt.Errorf("accessing relation %s in type %s: %w", e.relation, r.inType.name, ErrSymbolNotFound)
	}

	if e.subrelation == nil {
		return namespace.ComputedUserset(e.relation), nil
	}

	relationTypes, err := relation.DirectTypeReferences()
	if err != nil {
		return nil, err
	}

	for _, t := range relationTypes {
		if !t.IsResolved() {
			err := t.Resolve(r.inType)
			if err != nil {
				return nil, err
			}
		}

		var helperMethod func(tr *TypeReference) error

		helperMethod = func(tr *TypeReference) error {

			if tr.subRelation != "" {
				newRelation, ok := tr.instance.relations.Get(tr.subRelation)
				if !ok {
					return fmt.Errorf("accessing relation %s in type %s: %w", e.relation, r.inType.name, ErrSymbolNotFound)
				}

				newRelationTypes, err := newRelation.DirectTypeReferences()
				if err != nil {
					return err
				}

				for _, nt := range newRelationTypes {
					if !nt.IsResolved() {
						err := nt.Resolve(r.inType)
						if err != nil {
							return err
						}
					}

					err := helperMethod(nt)
					if err != nil {
						return err
					} else {
						return nil
					}
				}
			}

			subrelation, ok := tr.instance.relations.Get(*e.subrelation)
			if !ok {
				return fmt.Errorf("relation %s not found on type %s.%s %w", *e.subrelation, e.relation, *e.subrelation, ErrSymbolNotFound)
			}

			if !subrelation.VisibleTo(r.inType) {
				return ErrSymbolNotAccessible
			}

			return nil
		}

		err := helperMethod(t)
		if err != nil {
			return nil, err
		}
	}
	return namespace.TupleToUserset(relation.SpiceDBName(), *e.subrelation), nil
}

func (e *ReferenceRelationExpression) DirectTypeReferences(r *Relation) ([]*TypeReference, error) {
	return []*TypeReference{}, nil
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

func (e *SetRelationExpression) DirectTypeReferences(r *Relation) ([]*TypeReference, error) {
	leftTypes, err := e.left.DirectTypeReferences(r)
	if err != nil {
		return []*TypeReference{}, err
	}

	rightTypes, err := e.right.DirectTypeReferences(r)
	if err != nil {
		return []*TypeReference{}, err
	}

	return append(leftTypes, rightTypes...), nil
}
