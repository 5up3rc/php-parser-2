package name

import (
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/walker"
)

// FullyQualified node
type FullyQualified struct {
	NamespacedName string
	Parts          []node.Node
}

// NewFullyQualified node constuctor
func NewFullyQualified(Parts []node.Node, NamespacedName string) *FullyQualified {
	return &FullyQualified{
		NamespacedName,
		Parts,
	}
}

// Attributes returns node attributes as map
func (n *FullyQualified) Attributes() map[string]interface{} {
	return map[string]interface{}{
		"NamespacedName": n.NamespacedName,
	}
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *FullyQualified) Walk(v walker.Visitor) {
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
