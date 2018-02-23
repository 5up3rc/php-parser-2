// Package php7 parses PHP7
package php7

import (
	"github.com/z7zmey/php-parser/node/name"
	"io"

	"github.com/z7zmey/php-parser/comment"
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/node/stmt"
	"github.com/z7zmey/php-parser/position"
	"github.com/z7zmey/php-parser/token"
)

var rootnode node.Node
var comments comment.Comments
var positions position.Positions
var positionBuilder position.Builder
var namespace string

// Parse the php7 parser entrypoint
func Parse(src io.Reader, fName string) (node.Node, comment.Comments, position.Positions) {
	yyDebug = 0
	yyErrorVerbose = true
	rootnode = stmt.NewStmtList([]node.Node{}) //reset
	comments = comment.Comments{}
	positions = position.Positions{}
	positionBuilder = position.Builder{&positions}
	namespace = ""
	yyParse(newLexer(src, fName))
	return rootnode, comments, positions
}

// ListGetFirstNodeComments returns comments of a first node in the list
func ListGetFirstNodeComments(list []node.Node) []comment.Comment {
	if len(list) == 0 {
		return nil
	}

	node := list[0]

	return comments[node]
}

func joinNamespaceParts(list []node.Node) string {
	NSName := ""

	for _, part := range list {
		if NSName == "" {
			NSName = NSName + part.(*name.NamePart).Value
		} else {
			NSName = NSName + "\\" + part.(*name.NamePart).Value
		}
	}

	return NSName
}

type foreachVariable struct {
	node  node.Node
	byRef bool
}

type nodesWithEndToken struct {
	nodes    []node.Node
	endToken token.Token
}

type boolWithToken struct {
	value bool
	token *token.Token
}

type altSyntaxNode struct {
	node  node.Node
	isAlt bool
}
