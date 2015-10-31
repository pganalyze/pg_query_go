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
		`{"name": ["="], "lexpr": null, "rexpr": null}`,
		nodes.A_Expr{
			Kind: nodes.AEXPR_OP,
			Name: []nodes.Node{
				nodes.Value{
					Type: nodes.T_String,
					Str:  "=",
				},
			},
			Lexpr: nil,
			Rexpr: nil,
		},
	},
	{
		`{"name": ["="], "lexpr": {"COLUMNREF": {"fields": ["z"], "location": 22}}, "rexpr": {"A_CONST": {"val": 1, "location": 26}}, "location": 24}`,
		nodes.A_Expr{
			Kind: nodes.AEXPR_OP,
			Name: []nodes.Node{
				nodes.Value{
					Type: nodes.T_String,
					Str:  "=",
				},
			},
			Lexpr: nodes.ColumnRef{
				Fields: []nodes.Node{
					nodes.Value{
						Type: nodes.T_String,
						Str:  "z",
					},
				},
				Location: 22,
			},
			Rexpr: nodes.A_Const{
				Val: nodes.Value{
					Type: nodes.T_Integer,
					Ival: 1,
				},
				Location: 26,
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
