package semantic

type Relation struct {
	name       string
	visibility Visibility
	inType     *Type
	extensions []*ExtensionReference
	body       RelationExpression
}

func NewRelation(name string, visibility Visibility, body RelationExpression) (*Relation, error) {
	return &Relation{name: name, visibility: visibility, body: body}, nil
}

type RelationExpression interface {
	ToZanzibar() //Placeholder, return type TBD
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

func (e *SelfRelationExpression) ToZanzibar() {

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

func (e *ReferenceRelationExpression) ToZanzibar() {

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

func (e *SetRelationExpression) ToZanzibar() {

}
