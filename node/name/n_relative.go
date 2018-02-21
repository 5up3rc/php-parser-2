package name

import (
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/walker"
)

// Relative node
type Relative struct {
	NamespacedName string
	Parts          []node.Node
}

// NewRelative node constuctor
func NewRelative(Parts []node.Node, NamespacedName string) *Relative {
	return &Relative{
		NamespacedName,
		Parts,
	}
}

// Attributes returns node attributes as map
func (n *Relative) Attributes() map[string]interface{} {
	return map[string]interface{}{
		"NamespacedName": n.NamespacedName,
	}
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *Relative) Walk(v walker.Visitor) {
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
