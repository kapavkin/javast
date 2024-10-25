package javast

// A Visitor's Visit method is invoked for each node encountered by [Walk].
// If the result visitor v is not nil, [Walk] visits each of the children
// of node with the visitor v, followed by a call of v.Visit(nil).
type Visitor interface {
	Visit(node Node) (v Visitor)
}

// Walk traverses an AST in depth-first order: It starts by calling
// v.Visit(node); node must not be nil. If the visitor v returned by
// v.Visit(node) is not nil, Walk is invoked recursively with visitor
// v for each of the non-nil children of node, followed by a call of
// v.Visit(nil).
func Walk(v Visitor, node Node) {
	if v = v.Visit(node); v == nil {
		return
	}
	switch node.GetKind() {
	case ANNOTATED_TYPE:
		at := node.(AnnotatedTypeNode)
		for _, annotation := range at.GetAnnotations() {
			Walk(v, annotation)
		}
		Walk(v, at.GetUnderlyingType())
	default:
		panic("javast.Walk: unexpected node kind")
	}
}
