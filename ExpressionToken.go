package govaluate

/*
	Represents a single parsed token.
*/
type ExpressionToken struct {
	Kind     TokenKind
	FuncName string
	Value    interface{}
}
