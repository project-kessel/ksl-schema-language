package semantic

type Relation struct {
	visibility  Visibility
	cardinality Cardinality
	extensions  []*ExtensionReference
	body        *RelationReference
}

type SetExpression struct {
}
type RelationExpression struct {
	//The relation on this type to traverse. If nil, refers to 'this'
	via *RelationReference
	// The relation to traverse on referenced objects (ie: TupleToUserSet operations). If nil, refers directly to the referenced object (ie: ComputedUserSet)
	to *RelationReference
}
