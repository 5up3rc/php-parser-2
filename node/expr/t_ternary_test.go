package expr_test

import (
	"bytes"
	"testing"

	"github.com/z7zmey/php-parser/node/expr"

	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/node/stmt"
	"github.com/z7zmey/php-parser/php5"
	"github.com/z7zmey/php-parser/php7"
)

func TestTernary(t *testing.T) {
	src := `<? $a ? $b : $c;`

	expected := &stmt.StmtList{
		Stmts: []node.Node{
			&stmt.Expression{
				Expr: &expr.Ternary{
					Condition: &expr.Variable{VarName: &node.Identifier{Value: "$a"}},
					IfTrue:    &expr.Variable{VarName: &node.Identifier{Value: "$b"}},
					IfFalse:   &expr.Variable{VarName: &node.Identifier{Value: "$c"}},
				},
			},
		},
	}

	actual, _, _ := php7.Parse(bytes.NewBufferString(src), "test.php")
	assertEqual(t, expected, actual)

	actual, _, _ = php5.Parse(bytes.NewBufferString(src), "test.php")
	assertEqual(t, expected, actual)
}

func TestTernarySimple(t *testing.T) {
	src := `<? $a ? : $c;`

	expected := &stmt.StmtList{
		Stmts: []node.Node{
			&stmt.Expression{
				Expr: &expr.Ternary{
					Condition: &expr.Variable{VarName: &node.Identifier{Value: "$a"}},
					IfFalse:   &expr.Variable{VarName: &node.Identifier{Value: "$c"}},
				},
			},
		},
	}

	actual, _, _ := php7.Parse(bytes.NewBufferString(src), "test.php")
	assertEqual(t, expected, actual)

	actual, _, _ = php5.Parse(bytes.NewBufferString(src), "test.php")
	assertEqual(t, expected, actual)
}

func TestTernaryNestedTrue(t *testing.T) {
	src := `<? $a ? $b ? $c : $d : $e;`

	expected := &stmt.StmtList{
		Stmts: []node.Node{
			&stmt.Expression{
				Expr: &expr.Ternary{
					Condition: &expr.Variable{VarName: &node.Identifier{Value: "$a"}},
					IfTrue: &expr.Ternary{
						Condition: &expr.Variable{VarName: &node.Identifier{Value: "$b"}},
						IfTrue:    &expr.Variable{VarName: &node.Identifier{Value: "$c"}},
						IfFalse:   &expr.Variable{VarName: &node.Identifier{Value: "$d"}},
					},
					IfFalse: &expr.Variable{VarName: &node.Identifier{Value: "$e"}},
				},
			},
		},
	}

	actual, _, _ := php7.Parse(bytes.NewBufferString(src), "test.php")
	assertEqual(t, expected, actual)

	actual, _, _ = php5.Parse(bytes.NewBufferString(src), "test.php")
	assertEqual(t, expected, actual)
}

func TestTernaryNestedCond(t *testing.T) {
	src := `<? $a ? $b : $c ? $d : $e;`

	expected := &stmt.StmtList{
		Stmts: []node.Node{
			&stmt.Expression{
				Expr: &expr.Ternary{
					Condition: &expr.Ternary{
						Condition: &expr.Variable{VarName: &node.Identifier{Value: "$a"}},
						IfTrue:    &expr.Variable{VarName: &node.Identifier{Value: "$b"}},
						IfFalse:   &expr.Variable{VarName: &node.Identifier{Value: "$c"}},
					},
					IfTrue:  &expr.Variable{VarName: &node.Identifier{Value: "$d"}},
					IfFalse: &expr.Variable{VarName: &node.Identifier{Value: "$e"}},
				},
			},
		},
	}

	actual, _, _ := php7.Parse(bytes.NewBufferString(src), "test.php")
	assertEqual(t, expected, actual)

	actual, _, _ = php5.Parse(bytes.NewBufferString(src), "test.php")
	assertEqual(t, expected, actual)
}
