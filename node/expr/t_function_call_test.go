package expr_test

import (
	"bytes"
	"testing"

	"github.com/z7zmey/php-parser/node/name"

	"github.com/z7zmey/php-parser/node/expr"

	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/node/stmt"
	"github.com/z7zmey/php-parser/php5"
	"github.com/z7zmey/php-parser/php7"
)

func TestFunctionCall(t *testing.T) {
	src := `<? foo();`

	expected := &stmt.StmtList{
		Stmts: []node.Node{
			&stmt.Expression{
				Expr: &expr.FunctionCall{
					Function: &name.Name{NamespacedName: "",
						Parts: []node.Node{
							&name.NamePart{Value: "foo"},
						},
					},
					Arguments: []node.Node{},
				},
			},
		},
	}

	actual, _, _ := php7.Parse(bytes.NewBufferString(src), "test.php")
	assertEqual(t, expected, actual)

	actual, _, _ = php5.Parse(bytes.NewBufferString(src), "test.php")
	assertEqual(t, expected, actual)
}

func TestFunctionCallRelative(t *testing.T) {
	src := `<? namespace\foo();`

	expected := &stmt.StmtList{
		Stmts: []node.Node{
			&stmt.Expression{
				Expr: &expr.FunctionCall{
					Function: &name.Relative{NamespacedName: "",
						Parts: []node.Node{
							&name.NamePart{Value: "foo"},
						},
					},
					Arguments: []node.Node{},
				},
			},
		},
	}

	actual, _, _ := php7.Parse(bytes.NewBufferString(src), "test.php")
	assertEqual(t, expected, actual)

	actual, _, _ = php5.Parse(bytes.NewBufferString(src), "test.php")
	assertEqual(t, expected, actual)
}

func TestFunctionFullyQualified(t *testing.T) {
	src := `<? \foo([]);`

	expected := &stmt.StmtList{
		Stmts: []node.Node{
			&stmt.Expression{
				Expr: &expr.FunctionCall{
					Function: &name.FullyQualified{NamespacedName: "",
						Parts: []node.Node{
							&name.NamePart{Value: "foo"},
						},
					},
					Arguments: []node.Node{
						&node.Argument{
							Variadic:    false,
							IsReference: false,
							Expr: &expr.ShortArray{
								Items: []node.Node{},
							},
						},
					},
				},
			},
		},
	}

	actual, _, _ := php7.Parse(bytes.NewBufferString(src), "test.php")
	assertEqual(t, expected, actual)

	actual, _, _ = php5.Parse(bytes.NewBufferString(src), "test.php")
	assertEqual(t, expected, actual)
}

func TestFunctionCallVar(t *testing.T) {
	src := `<? $foo(yield $a);`

	expected := &stmt.StmtList{
		Stmts: []node.Node{
			&stmt.Expression{
				Expr: &expr.FunctionCall{
					Function: &expr.Variable{VarName: &node.Identifier{Value: "$foo"}},
					Arguments: []node.Node{
						&node.Argument{
							Variadic:    false,
							IsReference: false,
							Expr: &expr.Yield{
								Value: &expr.Variable{VarName: &node.Identifier{Value: "$a"}},
							},
						},
					},
				},
			},
		},
	}

	actual, _, _ := php7.Parse(bytes.NewBufferString(src), "test.php")
	assertEqual(t, expected, actual)

	actual, _, _ = php5.Parse(bytes.NewBufferString(src), "test.php")
	assertEqual(t, expected, actual)
}
