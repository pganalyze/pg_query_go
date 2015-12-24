package pg_query_test

import (
	"errors"
	"reflect"
	"testing"

	"github.com/kr/pretty"
	"github.com/lfittl/pg_query.go"
	nodes "github.com/lfittl/pg_query.go/nodes"
)

func strPtr(str string) *string {
	return &str
}

func strNode(str string) nodes.String {
	return nodes.String{Str: str}
}

func intNode(ival int64) nodes.Integer {
	return nodes.Integer{Ival: ival}
}

var parseTests = []struct {
	input        string
	expectedJSON string
	expectedTree pg_query.ParsetreeList
}{
	{
		"SELECT 1",
		`[{"SelectStmt": {"distinctClause": null, "intoClause": null, "targetList": [{"ResTarget": ` +
			`{"name": null, "indirection": null, "val": {"A_Const": {"val": {"Integer": {"ival": 1}}, "location": 7}}, "location": 7}}], ` +
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
								Val: nodes.Integer{
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
		`[{"SelectStmt": {"distinctClause": null, "intoClause": null, "targetList": [{"ResTarget": ` +
			`{"name": null, "indirection": null, "val": {"ColumnRef": {"fields": [{"A_Star": {}}], ` +
			`"location": 7}}, "location": 7}}], "fromClause": [{"RangeVar": {"schemaname": null, ` +
			`"relname": "x", "inhOpt": 2, "relpersistence": "p", "alias": null, "location": 14}}], ` +
			`"whereClause": {"A_Expr": {"kind": 0, "name": [{"String": {"str": "="}}], "lexpr": {"ColumnRef": {"fields": [{"String": {"str": "z"}}], ` +
			`"location": 22}}, "rexpr": {"A_Const": {"val": {"Integer": {"ival": 1}}, "location": 26}}, "location": 24}}, ` +
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
							strNode("="),
						},
						Lexpr: []nodes.Node{
							nodes.ColumnRef{
								Fields: []nodes.Node{
									strNode("z"),
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
			},
		},
	},
	{
		`INSERT INTO "schema_index_stats" ("snapshot_id","schema_index_id","size_bytes") VALUES (11710849,8448632,16384),(11710849,8448633,16384) RETURNING id`,
		`[{"InsertStmt": {"relation": {"RangeVar": {"schemaname": null, "relname": "schema_index_stats", ` +
			`"inhOpt": 2, "relpersistence": "p", "alias": null, "location": 12}}, "cols": [{"ResTarget": ` +
			`{"name": "snapshot_id", "indirection": null, "val": null, "location": 34}}, {"ResTarget": ` +
			`{"name": "schema_index_id", "indirection": null, "val": null, "location": 48}}, {"ResTarget": ` +
			`{"name": "size_bytes", "indirection": null, "val": null, "location": 66}}], "selectStmt": ` +
			`{"SelectStmt": {"distinctClause": null, "intoClause": null, "targetList": null, "fromClause": ` +
			`null, "whereClause": null, "groupClause": null, "havingClause": null, "windowClause": null, ` +
			`"valuesLists": [[{"A_Const": {"val": {"Integer": {"ival": 11710849}}, "location": 88}}, ` +
			`{"A_Const": {"val": {"Integer": {"ival": 8448632}}, "location": 97}}, {"A_Const": ` +
			`{"val": {"Integer": {"ival": 16384}}, "location": 105}}], [{"A_Const": {"val": ` +
			`{"Integer": {"ival": 11710849}}, "location": 113}}, {"A_Const": {"val": {"Integer": ` +
			`{"ival": 8448633}}, "location": 122}}, {"A_Const": {"val": {"Integer": {"ival": 16384}}, ` +
			`"location": 130}}]], "sortClause": null, "limitOffset": null, "limitCount": null, ` +
			`"lockingClause": null, "withClause": null, "op": 0, "all": false, "larg": null, ` +
			`"rarg": null}}, "returningList": [{"ResTarget": {"name": null, "indirection": null, ` +
			`"val": {"ColumnRef": {"fields": [{"String": {"str": "id"}}], "location": 147}}, ` +
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
									Val: nodes.Integer{
										Ival: 11710849,
									},
									Location: 88,
								},
								nodes.A_Const{
									Val: nodes.Integer{
										Ival: 8448632,
									},
									Location: 97,
								},
								nodes.A_Const{
									Val: nodes.Integer{
										Ival: 16384,
									},
									Location: 105,
								},
							},
							[]nodes.Node{
								nodes.A_Const{
									Val: nodes.Integer{
										Ival: 11710849,
									},
									Location: 113,
								},
								nodes.A_Const{
									Val: nodes.Integer{
										Ival: 8448633,
									},
									Location: 122,
								},
								nodes.A_Const{
									Val: nodes.Integer{
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
									strNode("id"),
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
	{
		"SELECT * FROM x WHERE y IN (?)",
		`[{"SelectStmt": {"distinctClause": null, "intoClause": null, "targetList": ` +
			`[{"ResTarget": {"name": null, "indirection": null, "val": {"ColumnRef": {"fields": ` +
			`[{"A_Star": {}}], "location": 7}}, "location": 7}}], "fromClause": [{"RangeVar": ` +
			`{"schemaname": null, "relname": "x", "inhOpt": 2, "relpersistence": "p", "alias": ` +
			`null, "location": 14}}], "whereClause": {"A_Expr": {"kind": 9, "name": [{"String": ` +
			`{"str": "="}}], "lexpr": {"ColumnRef": {"fields": [{"String": {"str": "y"}}], ` +
			`"location": 22}}, "rexpr": [{"ParamRef": {"number": 0, "location": 28}}], "location": ` +
			`24}}, "groupClause": null, "havingClause": null, "windowClause": null, "valuesLists": null, ` +
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
						Kind: nodes.AEXPR_IN,
						Name: []nodes.Node{
							strNode("="),
						},
						Lexpr: []nodes.Node{
							nodes.ColumnRef{
								Fields: []nodes.Node{
									strNode("y"),
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
			},
		},
	},
	{
		`SELECT n.nspname as "Schema",
  c.relname as "Name",
  CASE c.relkind WHEN 'r' THEN 'table' WHEN 'v' THEN 'view' WHEN 'm' THEN 'materialized view' WHEN 'i' THEN 'index' WHEN 'S' THEN 'sequence' WHEN 's' THEN 'special' WHEN 'f' THEN 'foreign table' END as "Type",
  pg_catalog.pg_get_userbyid(c.relowner) as "Owner"
FROM pg_catalog.pg_class c
     LEFT JOIN pg_catalog.pg_namespace n ON n.oid = c.relnamespace
WHERE c.relkind IN ('r','')
      AND n.nspname <> 'pg_catalog'
      AND n.nspname <> 'information_schema'
      AND n.nspname !~ '^pg_toast'
  AND pg_catalog.pg_table_is_visible(c.oid)
ORDER BY 1,2;`,
		`[{"SelectStmt": {"distinctClause": null, "intoClause": null, "targetList": ` +
			`[{"ResTarget": {"name": "Schema", "indirection": null, "val": {"ColumnRef": ` +
			`{"fields": [{"String": {"str": "n"}}, {"String": {"str": "nspname"}}], "location": 7}}, ` +
			`"location": 7}}, {"ResTarget": {"name": "Name", "indirection": null, "val": {"ColumnRef": ` +
			`{"fields": [{"String": {"str": "c"}}, {"String": {"str": "relname"}}], "location": 32}}, ` +
			`"location": 32}}, {"ResTarget": {"name": "Type", "indirection": null, "val": {"CaseExpr": ` +
			`{"casetype": 0, "casecollid": 0, "arg": {"ColumnRef": {"fields": [{"String": {"str": "c"}}, ` +
			`{"String": {"str": "relkind"}}], "location": 60}}, "args": [{"CaseWhen": {"expr": ` +
			`{"A_Const": {"val": {"String": {"str": "r"}}, "location": 75}}, "result": {"A_Const": ` +
			`{"val": {"String": {"str": "table"}}, "location": 84}}, "location": 70}}, {"CaseWhen": ` +
			`{"expr": {"A_Const": {"val": {"String": {"str": "v"}}, "location": 97}}, "result": ` +
			`{"A_Const": {"val": {"String": {"str": "view"}}, "location": 106}}, "location": 92}}, ` +
			`{"CaseWhen": {"expr": {"A_Const": {"val": {"String": {"str": "m"}}, "location": 118}}, ` +
			`"result": {"A_Const": {"val": {"String": {"str": "materialized view"}}, "location": 127}}, ` +
			`"location": 113}}, {"CaseWhen": {"expr": {"A_Const": {"val": {"String": {"str": "i"}}, ` +
			`"location": 152}}, "result": {"A_Const": {"val": {"String": {"str": "index"}}, ` +
			`"location": 161}}, "location": 147}}, {"CaseWhen": {"expr": {"A_Const": {"val": ` +
			`{"String": {"str": "S"}}, "location": 174}}, "result": {"A_Const": {"val": {"String": ` +
			`{"str": "sequence"}}, "location": 183}}, "location": 169}}, {"CaseWhen": {"expr": ` +
			`{"A_Const": {"val": {"String": {"str": "s"}}, "location": 199}}, "result": {"A_Const": ` +
			`{"val": {"String": {"str": "special"}}, "location": 208}}, "location": 194}}, {"CaseWhen": ` +
			`{"expr": {"A_Const": {"val": {"String": {"str": "f"}}, "location": 223}}, "result": ` +
			`{"A_Const": {"val": {"String": {"str": "foreign table"}}, "location": 232}}, ` +
			`"location": 218}}], "defresult": null, "location": 55}}, "location": 55}}, {"ResTarget": ` +
			`{"name": "Owner", "indirection": null, "val": {"FuncCall": {"funcname": [{"String": ` +
			`{"str": "pg_catalog"}}, {"String": {"str": "pg_get_userbyid"}}], "args": [{"ColumnRef": ` +
			`{"fields": [{"String": {"str": "c"}}, {"String": {"str": "relowner"}}], "location": 292}}], ` +
			`"agg_order": null, "agg_filter": null, "agg_within_group": false, "agg_star": false, ` +
			`"agg_distinct": false, "func_variadic": false, "over": null, "location": 265}}, ` +
			`"location": 265}}], "fromClause": [{"JoinExpr": {"jointype": 1, "isNatural": false, ` +
			`"larg": {"RangeVar": {"schemaname": "pg_catalog", "relname": "pg_class", "inhOpt": 2, ` +
			`"relpersistence": "p", "alias": {"Alias": {"aliasname": "c", "colnames": null}}, ` +
			`"location": 320}}, "rarg": {"RangeVar": {"schemaname": "pg_catalog", "relname": ` +
			`"pg_namespace", "inhOpt": 2, "relpersistence": "p", "alias": {"Alias": {"aliasname": ` +
			`"n", "colnames": null}}, "location": 357}}, "usingClause": null, "quals": {"A_Expr": ` +
			`{"kind": 0, "name": [{"String": {"str": "="}}], "lexpr": {"ColumnRef": {"fields": ` +
			`[{"String": {"str": "n"}}, {"String": {"str": "oid"}}], "location": 386}}, "rexpr": ` +
			`{"ColumnRef": {"fields": [{"String": {"str": "c"}}, {"String": {"str": "relnamespace"}}], ` +
			`"location": 394}}, "location": 392}}, "alias": null, "rtindex": 0}}], "whereClause": ` +
			`{"A_Expr": {"kind": 1, "name": null, "lexpr": {"A_Expr": {"kind": 1, "name": null, ` +
			`"lexpr": {"A_Expr": {"kind": 1, "name": null, "lexpr": {"A_Expr": {"kind": 1, "name": ` +
			`null, "lexpr": {"A_Expr": {"kind": 9, "name": [{"String": {"str": "="}}], "lexpr": ` +
			`{"ColumnRef": {"fields": [{"String": {"str": "c"}}, {"String": {"str": "relkind"}}], ` +
			`"location": 415}}, "rexpr": [{"A_Const": {"val": {"String": {"str": "r"}}, "location": ` +
			`429}}, {"A_Const": {"val": {"String": {"str": ""}}, "location": 433}}], "location": ` +
			`425}}, "rexpr": {"A_Expr": {"kind": 0, "name": [{"String": {"str": "<>"}}], "lexpr": ` +
			`{"ColumnRef": {"fields": [{"String": {"str": "n"}}, {"String": {"str": "nspname"}}], ` +
			`"location": 447}}, "rexpr": {"A_Const": {"val": {"String": {"str": "pg_catalog"}}, ` +
			`"location": 460}}, "location": 457}}, "location": 443}}, "rexpr": {"A_Expr": {"kind": 0, ` +
			`"name": [{"String": {"str": "<>"}}], "lexpr": {"ColumnRef": {"fields": [{"String": ` +
			`{"str": "n"}}, {"String": {"str": "nspname"}}], "location": 483}}, "rexpr": {"A_Const": ` +
			`{"val": {"String": {"str": "information_schema"}}, "location": 496}}, "location": 493}}, ` +
			`"location": 479}}, "rexpr": {"A_Expr": {"kind": 0, "name": [{"String": {"str": "!~"}}], ` +
			`"lexpr": {"ColumnRef": {"fields": [{"String": {"str": "n"}}, {"String": {"str": "nspname"}}], ` +
			`"location": 527}}, "rexpr": {"A_Const": {"val": {"String": {"str": "^pg_toast"}}, "location": ` +
			`540}}, "location": 537}}, "location": 523}}, "rexpr": {"FuncCall": {"funcname": [{"String": ` +
			`{"str": "pg_catalog"}}, {"String": {"str": "pg_table_is_visible"}}], "args": [{"ColumnRef": ` +
			`{"fields": [{"String": {"str": "c"}}, {"String": {"str": "oid"}}], "location": 589}}], ` +
			`"agg_order": null, "agg_filter": null, "agg_within_group": false, "agg_star": false, ` +
			`"agg_distinct": false, "func_variadic": false, "over": null, "location": 558}}, "location": ` +
			`554}}, "groupClause": null, "havingClause": null, "windowClause": null, "valuesLists": null, ` +
			`"sortClause": [{"SortBy": {"node": {"A_Const": {"val": {"Integer": {"ival": 1}}, "location": ` +
			`605}}, "sortby_dir": 0, "sortby_nulls": 0, "useOp": null, "location": -1}}, {"SortBy": ` +
			`{"node": {"A_Const": {"val": {"Integer": {"ival": 2}}, "location": 607}}, "sortby_dir": ` +
			`0, "sortby_nulls": 0, "useOp": null, "location": -1}}], "limitOffset": null, "limitCount": ` +
			`null, "lockingClause": null, "withClause": null, "op": 0, "all": false, "larg": null, "rarg": null}}]`,
		pg_query.ParsetreeList{
			Statements: []nodes.Node{
				nodes.SelectStmt{
					TargetList: []nodes.Node{
						nodes.ResTarget{
							Name: strPtr("Schema"),
							Val: nodes.ColumnRef{
								Fields: []nodes.Node{
									strNode("n"),
									strNode("nspname"),
								},
								Location: 7,
							},
							Location: 7,
						},
						nodes.ResTarget{
							Name: strPtr("Name"),
							Val: nodes.ColumnRef{
								Fields: []nodes.Node{
									strNode("c"),
									strNode("relname"),
								},
								Location: 32,
							},
							Location: 32,
						},
						nodes.ResTarget{
							Name: strPtr("Type"),
							Val: nodes.CaseExpr{
								Arg: nodes.ColumnRef{
									Fields: []nodes.Node{
										strNode("c"),
										strNode("relkind"),
									},
									Location: 60,
								},
								Args: []nodes.Node{
									nodes.CaseWhen{
										Expr:     nodes.A_Const{Val: strNode("r"), Location: 75},
										Result:   nodes.A_Const{Val: strNode("table"), Location: 84},
										Location: 70,
									},
									nodes.CaseWhen{
										Expr:     nodes.A_Const{Val: strNode("v"), Location: 97},
										Result:   nodes.A_Const{Val: strNode("view"), Location: 106},
										Location: 92,
									},
									nodes.CaseWhen{
										Expr:     nodes.A_Const{Val: strNode("m"), Location: 118},
										Result:   nodes.A_Const{Val: strNode("materialized view"), Location: 127},
										Location: 113,
									},
									nodes.CaseWhen{
										Expr:     nodes.A_Const{Val: strNode("i"), Location: 152},
										Result:   nodes.A_Const{Val: strNode("index"), Location: 161},
										Location: 147,
									},
									nodes.CaseWhen{
										Expr:     nodes.A_Const{Val: strNode("S"), Location: 174},
										Result:   nodes.A_Const{Val: strNode("sequence"), Location: 183},
										Location: 169,
									},
									nodes.CaseWhen{
										Expr:     nodes.A_Const{Val: strNode("s"), Location: 199},
										Result:   nodes.A_Const{Val: strNode("special"), Location: 208},
										Location: 194,
									},
									nodes.CaseWhen{
										Expr:     nodes.A_Const{Val: strNode("f"), Location: 223},
										Result:   nodes.A_Const{Val: strNode("foreign table"), Location: 232},
										Location: 218,
									},
								},
								Location: 55,
							},
							Location: 55,
						},
						nodes.ResTarget{
							Name: strPtr("Owner"),
							Val: nodes.FuncCall{
								Funcname: []nodes.Node{
									strNode("pg_catalog"),
									strNode("pg_get_userbyid"),
								},
								Args: []nodes.Node{
									nodes.ColumnRef{
										Fields: []nodes.Node{
											strNode("c"),
											strNode("relowner"),
										},
										Location: 292,
									},
								},
								Location: 265,
							},
							Location: 265,
						},
					},
					FromClause: []nodes.Node{
						nodes.JoinExpr{
							Jointype: 1,
							Larg: nodes.RangeVar{
								Schemaname:     strPtr("pg_catalog"),
								Relname:        strPtr("pg_class"),
								InhOpt:         2,
								Relpersistence: 'p',
								Alias: &nodes.Alias{
									Aliasname: strPtr("c"),
								},
								Location: 320,
							},
							Rarg: nodes.RangeVar{
								Schemaname:     strPtr("pg_catalog"),
								Relname:        strPtr("pg_namespace"),
								InhOpt:         2,
								Relpersistence: 'p',
								Alias: &nodes.Alias{
									Aliasname: strPtr("n"),
								},
								Location: 357,
							},
							Quals: nodes.A_Expr{
								Kind: 0,
								Name: []nodes.Node{strNode("=")},
								Lexpr: []nodes.Node{
									nodes.ColumnRef{
										Fields:   []nodes.Node{strNode("n"), strNode("oid")},
										Location: 386,
									},
								},
								Rexpr: []nodes.Node{
									nodes.ColumnRef{
										Fields:   []nodes.Node{strNode("c"), strNode("relnamespace")},
										Location: 394,
									},
								},
								Location: 392,
							},
						},
					},
					WhereClause: nodes.A_Expr{
						Kind: 1,
						Lexpr: []nodes.Node{
							nodes.A_Expr{
								Kind: 1,
								Lexpr: []nodes.Node{
									nodes.A_Expr{
										Kind: 1,
										Lexpr: []nodes.Node{
											nodes.A_Expr{
												Kind: 1,
												Lexpr: []nodes.Node{
													nodes.A_Expr{
														Kind: 9,
														Name: []nodes.Node{strNode("=")},
														Lexpr: []nodes.Node{
															nodes.ColumnRef{
																Fields:   []nodes.Node{strNode("c"), strNode("relkind")},
																Location: 415,
															},
														},
														Rexpr: []nodes.Node{
															nodes.A_Const{Val: strNode("r"), Location: 429},
															nodes.A_Const{Val: strNode(""), Location: 433},
														},
														Location: 425,
													},
												},
												Rexpr: []nodes.Node{
													nodes.A_Expr{
														Name: []nodes.Node{strNode("<>")},
														Lexpr: []nodes.Node{
															nodes.ColumnRef{
																Fields:   []nodes.Node{strNode("n"), strNode("nspname")},
																Location: 447,
															},
														},
														Rexpr: []nodes.Node{
															nodes.A_Const{Val: strNode("pg_catalog"), Location: 460},
														},
														Location: 457,
													},
												},
												Location: 443,
											},
										},
										Rexpr: []nodes.Node{
											nodes.A_Expr{
												Kind: 0,
												Name: []nodes.Node{strNode("<>")},
												Lexpr: []nodes.Node{
													nodes.ColumnRef{
														Fields:   []nodes.Node{strNode("n"), strNode("nspname")},
														Location: 483,
													},
												},
												Rexpr: []nodes.Node{
													nodes.A_Const{
														Val:      strNode("information_schema"),
														Location: 496,
													},
												},
												Location: 493,
											},
										},
										Location: 479,
									},
								},
								Rexpr: []nodes.Node{
									nodes.A_Expr{
										Kind: 0,
										Name: []nodes.Node{strNode("!~")},
										Lexpr: []nodes.Node{
											nodes.ColumnRef{
												Fields:   []nodes.Node{strNode("n"), strNode("nspname")},
												Location: 527,
											},
										},
										Rexpr: []nodes.Node{
											nodes.A_Const{
												Val:      strNode("^pg_toast"),
												Location: 540,
											},
										},
										Location: 537,
									},
								},
								Location: 523,
							},
						},
						Rexpr: []nodes.Node{
							nodes.FuncCall{
								Funcname: []nodes.Node{
									strNode("pg_catalog"),
									strNode("pg_table_is_visible"),
								},
								Args: []nodes.Node{
									nodes.ColumnRef{
										Fields:   []nodes.Node{strNode("c"), strNode("oid")},
										Location: 589,
									},
								},
								Location: 558,
							},
						},
						Location: 554,
					},
					SortClause: []nodes.Node{
						nodes.SortBy{
							Node: nodes.A_Const{
								Val:      intNode(1),
								Location: 605,
							},
							Location: -1,
						},
						nodes.SortBy{
							Node: nodes.A_Const{
								Val:      intNode(2),
								Location: 607,
							},
							Location: -1,
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
			t.Errorf("Unmarshal(%s)\ndiff %s\n\n", actualJSON, pretty.Diff(test.expectedTree, actualTree))
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
