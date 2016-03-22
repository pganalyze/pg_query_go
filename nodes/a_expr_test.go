package pg_query_test

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/kr/pretty"
	nodes "github.com/lfittl/pg_query_go/nodes"
	"github.com/lfittl/pg_query_go/util"
)

var aExprTests = []struct {
	jsonText     string
	expectedNode nodes.A_Expr
}{
	{
		`{"name": [{"String": {"str": "="}}], "lexpr": null, "rexpr": null}`,
		nodes.A_Expr{
			Kind: nodes.AEXPR_OP,
			Name: util.MakeListNode([]nodes.Node{
				util.MakeStrNode("="),
			}),
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
			Name: util.MakeListNode([]nodes.Node{
				util.MakeStrNode("="),
			}),
			Lexpr: nodes.ColumnRef{
				Fields: util.MakeListNode([]nodes.Node{
					util.MakeStrNode("z"),
				}),
				Location: 22,
			},
			Rexpr: nodes.A_Const{
				Val: nodes.Integer{
					Ival: 1,
				},
				Location: 26,
			},
			Location: 24,
		},
	},
	{
		`{"kind": 6, "name": [{"String": {"str": "="}}], "lexpr": {"ColumnRef": ` +
			`{"fields": [{"String": {"str": "y"}}], "location": 22}}, "rexpr": {"List": ` +
			`{"items": [{"ParamRef": {"number": 0, "location": 28}}]}}, "location": 24}`,
		nodes.A_Expr{
			Kind: nodes.AEXPR_IN,
			Name: util.MakeListNode([]nodes.Node{
				util.MakeStrNode("="),
			}),
			Lexpr: nodes.ColumnRef{
				Fields: util.MakeListNode([]nodes.Node{
					util.MakeStrNode("y"),
				}),
				Location: 22,
			},
			Rexpr: nodes.List{
				Items: []nodes.Node{
					nodes.ParamRef{
						Number:   0,
						Location: 28,
					},
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
			t.Errorf("Unmarshal(%s)\ndiff %s\n\n", test.jsonText, pretty.Diff(test.expectedNode, actualTree))
		}
	}
}
