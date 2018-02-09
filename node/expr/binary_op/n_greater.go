package binary_op

import (
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/walker"
)

// Greater node
type Greater struct {
	Left  node.Node
	Right node.Node
}

// NewGreater node constuctor
func NewGreater(Variable node.Node, Expression node.Node) *Greater {
	return &Greater{
		Variable,
		Expression,
	}
}

// Attributes returns node attributes as map
func (n *Greater) Attributes() map[string]interface{} {
	return nil
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *Greater) Walk(v walker.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Left != nil {
		vv := v.GetChildrenVisitor("Left")
		n.Left.Walk(vv)
	}

	if n.Right != nil {
		vv := v.GetChildrenVisitor("Right")
		n.Right.Walk(vv)
	}

	v.LeaveNode(n)
}
