package semantic

type Name interface {
	String() string
}

type LiteralName string

func (n LiteralName) String() string {
	return string(n)
}
