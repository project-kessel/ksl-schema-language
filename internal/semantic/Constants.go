package semantic

type Visibility struct {
	NamespaceVisible bool
	SchemaVisible    bool
}

var (
	VisibilityPrivate = Visibility{
		NamespaceVisible: false,
		SchemaVisible:    false,
	}
	VisibilityInternal = Visibility{
		NamespaceVisible: true,
		SchemaVisible:    false,
	}
	VisibilityPublic = Visibility{
		NamespaceVisible: true,
		SchemaVisible:    true,
	}
)

type Cardinality int

const (
	CardinalityAtMostOne Cardinality = iota
	CardinalityExactlyOne
	CardinalityAtLeastOne
	CardinalityAny
)

type SetOperation int

const (
	SetOperationUnion SetOperation = iota
	SetOperationIntersection
	SetOperationExclusion
)
