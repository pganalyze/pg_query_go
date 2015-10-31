package pg_query_test

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/lfittl/pg_query.go"
	nodes "github.com/lfittl/pg_query.go/nodes"
)

func strPtr(str string) *string {
	return &str
}

var queryTests = []struct {
	input        string
	expectedJson string
	expectedTree pg_query.ParsetreeList
}{
	{
		"SELECT 1",
		`[{"SELECT": {"distinctClause": null, "intoClause": null, "targetList": [{"RESTARGET": ` +
			`{"name": null, "indirection": null, "val": {"A_CONST": {"val": 1, "location": 7}}, "location": 7}}], ` +
			`"fromClause": null, "whereClause": null, "groupClause": null, "havingClause": null, ` +
			`"windowClause": null, "valuesLists": null, "sortClause": null, "limitOffset": null, ` +
			`"limitCount": null, "lockingClause": null, "withClause": null, "op": 0, "all": false, ` +
			`"larg": null, "rarg": null}}]`,
		pg_query.ParsetreeList{
			Statements: []nodes.Node{
				nodes.SelectStmt{
					TargetList: []nodes.Node{
						nodes.ResTarget{
							Val: nodes.A_Const{
								Val: nodes.Value{
									Type: nodes.T_Integer,
									Ival: 1,
								},
								Location: 7,
							},
							Location: 7,
						},
					},
				},
			},
		},
	},
	{
		"SELECT * FROM x WHERE z = 1",
		`[{"SELECT": {"distinctClause": null, "intoClause": null, "targetList": [{"RESTARGET": ` +
			`{"name": null, "indirection": null, "val": {"COLUMNREF": {"fields": [{"A_STAR": {}}], ` +
			`"location": 7}}, "location": 7}}], "fromClause": [{"RANGEVAR": {"schemaname": null, ` +
			`"relname": "x", "inhOpt": 2, "relpersistence": "p", "alias": null, "location": 14}}], ` +
			`"whereClause": {"AEXPR": {"name": ["="], "lexpr": {"COLUMNREF": {"fields": ["z"], ` +
			`"location": 22}}, "rexpr": {"A_CONST": {"val": 1, "location": 26}}, "location": 24}}, ` +
			`"groupClause": null, "havingClause": null, "windowClause": null, "valuesLists": null, ` +
			`"sortClause": null, "limitOffset": null, "limitCount": null, "lockingClause": null, ` +
			`"withClause": null, "op": 0, "all": false, "larg": null, "rarg": null}}]`,
		pg_query.ParsetreeList{
			Statements: []nodes.Node{
				nodes.SelectStmt{
					TargetList: []nodes.Node{
						nodes.ResTarget{
							Val: nodes.ColumnRef{
								Fields: []nodes.Node{
									nodes.A_Star{},
								},
								Location: 7,
							},
							Location: 7,
						},
					},
					FromClause: []nodes.Node{
						nodes.RangeVar{
							Relname:        strPtr("x"),
							InhOpt:         nodes.INH_DEFAULT,
							Relpersistence: 'p',
							Location:       14,
						},
					},
					WhereClause: nodes.A_Expr{
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
			},
		},
	},
}

func TestParse(t *testing.T) {
	for _, test := range queryTests {
		actualJson := pg_query.Parse(test.input)
		if actualJson != test.expectedJson {
			t.Errorf("Parse(%s)\nexpected %s\nactual %s\n\n", test.input, test.expectedJson, actualJson)
		}

		var actualTree pg_query.ParsetreeList
		err := json.Unmarshal([]byte(actualJson), &actualTree)

		if err != nil {
			t.Errorf("Unmarshal(%s)\nerror %s\n\n", actualJson, err)
		} else if !reflect.DeepEqual(actualTree, test.expectedTree) {
			t.Errorf("Unmarshal(%s)\nexpected %s\nactual %s\n\n", actualJson, test.expectedTree, actualTree)
		}
	}
}
