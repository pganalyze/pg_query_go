package pg_query_test

import (
	"errors"
	"reflect"
	"testing"

	"github.com/lfittl/pg_query.go"
	nodes "github.com/lfittl/pg_query.go/nodes"
)

func strPtr(str string) *string {
	return &str
}

var parseTests = []struct {
	input        string
	expectedJSON string
	expectedTree pg_query.ParsetreeList
}{
	{
		"SELECT 1",
		`[{"SELECT": {"distinctClause": null, "intoClause": null, "targetList": [{"RESTARGET": ` +
			`{"name": null, "indirection": null, "val": {"A_CONST": {"type": "integer", "val": 1, "location": 7}}, "location": 7}}], ` +
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
								Type: "integer",
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
			`"location": 22}}, "rexpr": {"A_CONST": {"type": "integer", "val": 1, "location": 26}}, "location": 24}}, ` +
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
							Type: "integer",
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
	{
		`INSERT INTO "schema_index_stats" ("snapshot_id","schema_index_id","size_bytes") VALUES (11710849,8448632,16384),(11710849,8448633,16384) RETURNING id`,
		`[{"INSERT INTO": {"relation": {"RANGEVAR": {"schemaname": null, "relname": ` +
			`"schema_index_stats", "inhOpt": 2, "relpersistence": "p", "alias": null, ` +
			`"location": 12}}, "cols": [{"RESTARGET": {"name": "snapshot_id", "indirection": ` +
			`null, "val": null, "location": 34}}, {"RESTARGET": {"name": "schema_index_id", ` +
			`"indirection": null, "val": null, "location": 48}}, {"RESTARGET": {"name": ` +
			`"size_bytes", "indirection": null, "val": null, "location": 66}}], "selectStmt": ` +
			`{"SELECT": {"distinctClause": null, "intoClause": null, "targetList": null, ` +
			`"fromClause": null, "whereClause": null, "groupClause": null, "havingClause": ` +
			`null, "windowClause": null, "valuesLists": [[{"A_CONST": {"type": "integer", "val": ` +
			`11710849, "location": 88}}, {"A_CONST": {"type": "integer", "val": 8448632, "location": 97}}, ` +
			`{"A_CONST": {"type": "integer", "val": 16384, "location": 105}}], [{"A_CONST": ` +
			`{"type": "integer", "val": 11710849, "location": 113}}, {"A_CONST": {"type": ` +
			`"integer", "val": 8448633, "location": 122}}, {"A_CONST": {"type": "integer", ` +
			`"val": 16384, "location": 130}}]], "sortClause": null, "limitOffset": null, ` +
			`"limitCount": null, "lockingClause": null, "withClause": null, "op": 0, "all": ` +
			`false, "larg": null, "rarg": null}}, "returningList": [{"RESTARGET": {"name": null, ` +
			`"indirection": null, "val": {"COLUMNREF": {"fields": ["id"], "location": 147}}, ` +
			`"location": 147}}], "withClause": null}}]`,
		pg_query.ParsetreeList{
			Statements: []nodes.Node{
				nodes.InsertStmt{
					Relation: &nodes.RangeVar{
						Relname:        strPtr("schema_index_stats"),
						InhOpt:         nodes.INH_DEFAULT,
						Relpersistence: []byte("p")[0],
						Location:       12,
					},
					Cols: []nodes.Node{
						nodes.ResTarget{
							Name:     strPtr("snapshot_id"),
							Location: 34,
						},
						nodes.ResTarget{
							Name:     strPtr("schema_index_id"),
							Location: 48,
						},
						nodes.ResTarget{
							Name:     strPtr("size_bytes"),
							Location: 66,
						},
					},
					SelectStmt: nodes.SelectStmt{
						ValuesLists: [][]nodes.Node{
							[]nodes.Node{
								nodes.A_Const{
									Type: "integer",
									Val: nodes.Value{
										Type: nodes.T_Integer,
										Ival: 11710849,
									},
									Location: 88,
								},
								nodes.A_Const{
									Type: "integer",
									Val: nodes.Value{
										Type: nodes.T_Integer,
										Ival: 8448632,
									},
									Location: 97,
								},
								nodes.A_Const{
									Type: "integer",
									Val: nodes.Value{
										Type: nodes.T_Integer,
										Ival: 16384,
									},
									Location: 105,
								},
							},
							[]nodes.Node{
								nodes.A_Const{
									Type: "integer",
									Val: nodes.Value{
										Type: nodes.T_Integer,
										Ival: 11710849,
									},
									Location: 113,
								},
								nodes.A_Const{
									Type: "integer",
									Val: nodes.Value{
										Type: nodes.T_Integer,
										Ival: 8448633,
									},
									Location: 122,
								},
								nodes.A_Const{
									Type: "integer",
									Val: nodes.Value{
										Type: nodes.T_Integer,
										Ival: 16384,
									},
									Location: 130,
								},
							},
						},
					},
					ReturningList: []nodes.Node{
						nodes.ResTarget{
							Val: nodes.ColumnRef{
								Fields: []nodes.Node{
									nodes.Value{
										Type: nodes.T_String,
										Str:  "id",
									},
								},
								Location: 147,
							},
							Location: 147,
						},
					},
				},
			},
		},
	},
}

func TestParse(t *testing.T) {
	for _, test := range parseTests {
		actualJSON, err := pg_query.ParseToJSON(test.input)
		if err != nil {
			t.Errorf("Parse(%s)\nerror %s\n\n", test.input, err)
		} else if actualJSON != test.expectedJSON {
			t.Errorf("Parse(%s)\nexpected %s\nactual %s\n\n", test.input, test.expectedJSON, actualJSON)
		}

		actualTree, err := pg_query.Parse(test.input)

		if err != nil {
			t.Errorf("Unmarshal(%s)\nerror %s\n\n", actualJSON, err)
		} else if !reflect.DeepEqual(actualTree, test.expectedTree) {
			t.Errorf("Unmarshal(%s)\nexpected %s\nactual %s\n\n", actualJSON, test.expectedTree, actualTree)
		}
	}
}

var parseErrorTests = []struct {
	input       string
	expectedErr error
}{
	{
		"SELECT $",
		errors.New("syntax error at or near \"$\""),
	},
	{
		"SELECT * FROM y WHERE x IN (?, ",
		errors.New("syntax error at end of input"),
	},
}

func TestParseError(t *testing.T) {
	for _, test := range parseErrorTests {
		_, actualErr := pg_query.Parse(test.input)

		if actualErr == nil {
			t.Errorf("Parse(%s)\nexpected error but none returned\n\n", test.input)
		} else if !reflect.DeepEqual(actualErr, test.expectedErr) {
			t.Errorf("Parse(%s)\nexpected error %s\nactual error %s\n\n", test.input, test.expectedErr, actualErr)
		}
	}
}
