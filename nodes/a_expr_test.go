package pg_query_test

import (
	"encoding/json"
	"reflect"
	"testing"

	nodes "github.com/lfittl/pg_query.go/nodes"
)

var aExprTests = []struct {
	jsonText     string
	expectedNode nodes.A_Expr
}{
	{
		`{"name": [{"String": {"str": "="}}], "lexpr": null, "rexpr": null}`,
		nodes.A_Expr{
			Kind: nodes.AEXPR_OP,
			Name: []nodes.Node{
				nodes.String{
					Str: "=",
				},
			},
			Lexpr: nil,
			Rexpr: nil,
		},
	},
	{
		`{"name": [{"String": {"str": "="}}], "lexpr": {"ColumnRef": {"fields": ` +
			`[{"String": {"str": "z"}}], "location": 22}}, "rexpr": {"A_Const": {"val": ` +
			`{"Integer": {"ival": 1}}, "location": 26}}, "location": 24}`,
		nodes.A_Expr{
			Kind: nodes.AEXPR_OP,
			Name: []nodes.Node{
				nodes.String{
					Str: "=",
				},
			},
			Lexpr: []nodes.Node{
				nodes.ColumnRef{
					Fields: []nodes.Node{
						nodes.String{
							Str: "z",
						},
					},
					Location: 22,
				},
			},
			Rexpr: []nodes.Node{
				nodes.A_Const{
					Val: nodes.Integer{
						Ival: 1,
					},
					Location: 26,
				},
			},
			Location: 24,
		},
	},
	{
		`{"kind": 9, "name": [{"String": {"str": "="}}], "lexpr": {"ColumnRef": ` +
			`{"fields": [{"String": {"str": "y"}}], "location": 22}}, "rexpr": [{"ParamRef": ` +
			`{"number": 0, "location": 28}}], "location": 24}`,
		nodes.A_Expr{
			Kind: nodes.AEXPR_IN,
			Name: []nodes.Node{
				nodes.String{
					Str: "=",
				},
			},
			Lexpr: []nodes.Node{
				nodes.ColumnRef{
					Fields: []nodes.Node{
						nodes.String{
							Str: "y",
						},
					},
					Location: 22,
				},
			},
			Rexpr: []nodes.Node{
				nodes.ParamRef{
					Number:   0,
					Location: 28,
				},
			},
			Location: 24,
		},
	},
}

func TestAExpr(t *testing.T) {
	for _, test := range aExprTests {
		var actualTree nodes.A_Expr
		err := json.Unmarshal([]byte(test.jsonText), &actualTree)

		if err != nil {
			t.Errorf("Unmarshal(%s)\nerror %s\n\n", test.jsonText, err)
		} else if !reflect.DeepEqual(actualTree, test.expectedNode) {
			t.Errorf("Unmarshal(%s)\nexpected %s\nactual %s\n\n", test.jsonText, test.expectedNode, actualTree)
		}
	}
}
