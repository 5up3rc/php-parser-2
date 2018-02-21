package name

import (
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/walker"
)

// Name node
type Name struct {
	NamespacedName string
	Parts          []node.Node
}

// NewName node constuctor
func NewName(Parts []node.Node, NamespacedName string) *Name {
	return &Name{
		NamespacedName,
		Parts,
	}
}

// Attributes returns node attributes as map
func (n *Name) Attributes() map[string]interface{} {
	return map[string]interface{}{
		"NamespacedName": n.NamespacedName,
	}
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *Name) Walk(v walker.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Parts != nil {
		vv := v.GetChildrenVisitor("Parts")
		for _, nn := range n.Parts {
			if nn != nil {
				nn.Walk(vv)
			}
		}
	}

	v.LeaveNode(n)
}
