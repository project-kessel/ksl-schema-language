package semantic

type Extension struct {
	name       string
	visibility Visibility
	types      []*DynamicType
	namespace  *Namespace
}

func NewExtension(name string, visibility Visibility, types []*DynamicType) *Extension {
	return &Extension{name: name, visibility: visibility, types: types}
}

func (e *Extension) Apply(fromNS *Namespace, fromType *Type, fromRelation *Relation, params map[string]string) error {
	params["NAMESPACE"] = fromNS.name
	if fromType != nil {
		params["TYPE"] = fromType.name

		if fromRelation != nil {
			params["RELATION"] = fromRelation.name
		}
	}

	for _, t := range e.types {
		generated, err := t.ToType(e.namespace, params)
		if err != nil {
			return err
		}

		existing, found := e.namespace.types.Get(generated.name)
		if !found {
			err = e.namespace.AddType(generated)
			if err != nil {
				return err
			}
		} else {
			err = generated.relations.Iterate(func(s string, relation *Relation) error {
				if _, ok := existing.relations.Get(relation.name); ok && relation.generatedFrom.IgnoreDuplicates {
					return nil
				}
				return existing.AddRelation(relation)
			})
			if err != nil {
				return err
			}
		}
	}

	return nil
}

type DynamicType struct {
	Name       Name
	Visibility Visibility
	Relations  []*DynamicRelation
}

func (dt *DynamicType) ToType(m *Namespace, params map[string]string) (*Type, error) {
	t := NewType(dt.Name.String(params), m, dt.Visibility)

	for _, dr := range dt.Relations {
		r, err := dr.ToRelation(t, params)
		if err != nil {
			return nil, err
		}

		err = t.AddRelation(r)
		if err != nil {
			return nil, err
		}
	}
	return t, nil
}

type DynamicRelation struct {
	Name             Name
	Visibility       Visibility
	IgnoreDuplicates bool
	Body             DynamicRelationBody
}

func (dr *DynamicRelation) ToRelation(t *Type, params map[string]string) (*Relation, error) {
	body, err := dr.Body.ToRelationExpression(params)
	if err != nil {
		return nil, err
	}

	return NewRelation(dr.Name.String(params), t, dr.Visibility, body, dr)
}

type DynamicRelationBody interface {
	ToRelationExpression(params map[string]string) (RelationExpression, error)
}

func (e *SelfRelationExpression) ToRelationExpression(params map[string]string) (RelationExpression, error) {
	return e, nil
}

type DynamicReferenceRelationExpression struct {
	Relation    Name
	SubRelation Name
}

func (e *DynamicReferenceRelationExpression) ToRelationExpression(params map[string]string) (RelationExpression, error) {
	var subrelation *string
	if e.SubRelation != nil {
		r := e.SubRelation.String(params)
		subrelation = &r
	} else {
		subrelation = nil
	}

	return NewReferenceRelationExpression(e.Relation.String(params), subrelation), nil
}

type DynamicSetRelationExpression struct {
	Kind  string
	Left  DynamicRelationBody
	Right DynamicRelationBody
}

func (e DynamicSetRelationExpression) ToRelationExpression(params map[string]string) (RelationExpression, error) {
	left, err := e.Left.ToRelationExpression(params)
	if err != nil {
		return nil, err
	}

	right, err := e.Right.ToRelationExpression(params)
	if err != nil {
		return nil, err
	}

	return NewSetRelationExpression(e.Kind, left, right), nil
}
