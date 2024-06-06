package semantic

type Visibility struct {
	ModuleVisible bool
	SchemaVisible bool
}

var (
	VisibilityPrivate = Visibility{
		ModuleVisible: false,
		SchemaVisible: false,
	}
	VisibilityInternal = Visibility{
		ModuleVisible: true,
		SchemaVisible: false,
	}
	VisibilityPublic = Visibility{
		ModuleVisible: true,
		SchemaVisible: true,
	}
)

type Cardinality int

const (
	CardinalityZeroOrOne Cardinality = iota
	CardinalityExactlyOne
	CardinalityAtLeastOne
	CardinalityZeroToMany
)

type SetOperation int

const (
	SetOperationUnion SetOperation = iota
	SetOperationIntersection
	SetOperationExclusion
)
