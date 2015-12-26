package pg_query_test

import (
	"errors"
	"reflect"
	"testing"

	"github.com/kr/pretty"
	"github.com/lfittl/pg_query.go"
	nodes "github.com/lfittl/pg_query.go/nodes"
	util "github.com/lfittl/pg_query.go/util"
)

var parseTests = []struct {
	input        string
	expectedJSON string
	expectedTree pg_query.ParsetreeList
}{
	{
		"SELECT 1",
		`[{"SelectStmt": {"targetList": [{"ResTarget": {"val": {"A_Const": {"val": ` +
			`{"Integer": {"ival": 1}}, "location": 7}}, "location": 7}}], "op": 0}}]`,
		pg_query.ParsetreeList{
			Statements: []nodes.Node{
				nodes.SelectStmt{
					TargetList: util.MakeListNode([]nodes.Node{
						nodes.ResTarget{
							Val: nodes.A_Const{
								Val:      util.MakeIntNode(1),
								Location: 7,
							},
							Location: 7,
						},
					}),
				},
			},
		},
	},
	{
		"SELECT * FROM x WHERE z = 1",
		`[{"SelectStmt": {"targetList": [{"ResTarget": {"val": {"ColumnRef": {"fields": ` +
			`[{"A_Star": {}}], "location": 7}}, "location": 7}}], "fromClause": [{"RangeVar": ` +
			`{"relname": "x", "inhOpt": 2, "relpersistence": "p", "location": 14}}], "whereClause": ` +
			`{"A_Expr": {"kind": 0, "name": [{"String": {"str": "="}}], "lexpr": {"ColumnRef": ` +
			`{"fields": [{"String": {"str": "z"}}], "location": 22}}, "rexpr": {"A_Const": ` +
			`{"val": {"Integer": {"ival": 1}}, "location": 26}}, "location": 24}}, "op": 0}}]`,
		pg_query.ParsetreeList{
			Statements: []nodes.Node{
				nodes.SelectStmt{
					TargetList: util.MakeListNode([]nodes.Node{
						nodes.ResTarget{
							Val: nodes.ColumnRef{
								Fields: util.MakeListNode([]nodes.Node{
									nodes.A_Star{},
								}),
								Location: 7,
							},
							Location: 7,
						},
					}),
					FromClause: util.MakeListNode([]nodes.Node{
						nodes.RangeVar{
							Relname:        util.MakeStrPtr("x"),
							InhOpt:         nodes.INH_DEFAULT,
							Relpersistence: 'p',
							Location:       14,
						},
					}),
					WhereClause: nodes.A_Expr{
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
							Val:      util.MakeIntNode(1),
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
		`[{"InsertStmt": {"relation": {"RangeVar": {"relname": "schema_index_stats", ` +
			`"inhOpt": 2, "relpersistence": "p", "location": 12}}, "cols": [{"ResTarget": ` +
			`{"name": "snapshot_id", "location": 34}}, {"ResTarget": {"name": "schema_index_id", ` +
			`"location": 48}}, {"ResTarget": {"name": "size_bytes", "location": 66}}], "selectStmt": ` +
			`{"SelectStmt": {"valuesLists": [[{"A_Const": {"val": {"Integer": {"ival": 11710849}}, ` +
			`"location": 88}}, {"A_Const": {"val": {"Integer": {"ival": 8448632}}, "location": 97}}, ` +
			`{"A_Const": {"val": {"Integer": {"ival": 16384}}, "location": 105}}], [{"A_Const": {"val": ` +
			`{"Integer": {"ival": 11710849}}, "location": 113}}, {"A_Const": {"val": {"Integer": ` +
			`{"ival": 8448633}}, "location": 122}}, {"A_Const": {"val": {"Integer": {"ival": 16384}}, ` +
			`"location": 130}}]], "op": 0}}, "returningList": [{"ResTarget": {"val": ` +
			`{"ColumnRef": {"fields": [{"String": {"str": "id"}}], "location": 147}}, "location": 147}}]}}]`,
		pg_query.ParsetreeList{
			Statements: []nodes.Node{
				nodes.InsertStmt{
					Relation: &nodes.RangeVar{
						Relname:        util.MakeStrPtr("schema_index_stats"),
						InhOpt:         nodes.INH_DEFAULT,
						Relpersistence: 'p',
						Location:       12,
					},
					Cols: util.MakeListNode([]nodes.Node{
						nodes.ResTarget{
							Name:     util.MakeStrPtr("snapshot_id"),
							Location: 34,
						},
						nodes.ResTarget{
							Name:     util.MakeStrPtr("schema_index_id"),
							Location: 48,
						},
						nodes.ResTarget{
							Name:     util.MakeStrPtr("size_bytes"),
							Location: 66,
						},
					}),
					SelectStmt: nodes.SelectStmt{
						ValuesLists: [][]nodes.Node{
							[]nodes.Node{
								nodes.A_Const{
									Val:      util.MakeIntNode(11710849),
									Location: 88,
								},
								nodes.A_Const{
									Val:      util.MakeIntNode(8448632),
									Location: 97,
								},
								nodes.A_Const{
									Val:      util.MakeIntNode(16384),
									Location: 105,
								},
							},
							[]nodes.Node{
								nodes.A_Const{
									Val:      util.MakeIntNode(11710849),
									Location: 113,
								},
								nodes.A_Const{
									Val:      util.MakeIntNode(8448633),
									Location: 122,
								},
								nodes.A_Const{
									Val:      util.MakeIntNode(16384),
									Location: 130,
								},
							},
						},
					},
					ReturningList: util.MakeListNode([]nodes.Node{
						nodes.ResTarget{
							Val: nodes.ColumnRef{
								Fields: util.MakeListNode([]nodes.Node{
									util.MakeStrNode("id"),
								}),
								Location: 147,
							},
							Location: 147,
						},
					}),
				},
			},
		},
	},
	{
		"SELECT * FROM x WHERE y IN (?)",
		`[{"SelectStmt": {"targetList": [{"ResTarget": {"val": {"ColumnRef": {"fields": ` +
			`[{"A_Star": {}}], "location": 7}}, "location": 7}}], "fromClause": [{"RangeVar": ` +
			`{"relname": "x", "inhOpt": 2, "relpersistence": "p", "location": 14}}], "whereClause": ` +
			`{"A_Expr": {"kind": 9, "name": [{"String": {"str": "="}}], "lexpr": {"ColumnRef": ` +
			`{"fields": [{"String": {"str": "y"}}], "location": 22}}, "rexpr": [{"ParamRef": ` +
			`{"location": 28}}], "location": 24}}, "op": 0}}]`,
		pg_query.ParsetreeList{
			Statements: []nodes.Node{
				nodes.SelectStmt{
					TargetList: util.MakeListNode([]nodes.Node{
						nodes.ResTarget{
							Val: nodes.ColumnRef{
								Fields: util.MakeListNode([]nodes.Node{
									nodes.A_Star{},
								}),
								Location: 7,
							},
							Location: 7,
						},
					}),
					FromClause: util.MakeListNode([]nodes.Node{
						nodes.RangeVar{
							Relname:        util.MakeStrPtr("x"),
							InhOpt:         nodes.INH_DEFAULT,
							Relpersistence: 'p',
							Location:       14,
						},
					}),
					WhereClause: nodes.A_Expr{
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
						Rexpr: util.MakeListNode([]nodes.Node{
							nodes.ParamRef{
								Number:   0,
								Location: 28,
							},
						}),
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
		`[{"SelectStmt": {"targetList": [{"ResTarget": {"name": "Schema", "val": {"ColumnRef": ` +
			`{"fields": [{"String": {"str": "n"}}, {"String": {"str": "nspname"}}], "location": 7}}, ` +
			`"location": 7}}, {"ResTarget": {"name": "Name", "val": {"ColumnRef": {"fields": [{"String": ` +
			`{"str": "c"}}, {"String": {"str": "relname"}}], "location": 32}}, "location": 32}}, ` +
			`{"ResTarget": {"name": "Type", "val": {"CaseExpr": {` +
			`"arg": {"ColumnRef": {"fields": [{"String": {"str": "c"}}, {"String": {"str": "relkind"}}], ` +
			`"location": 60}}, "args": [{"CaseWhen": {"expr": {"A_Const": {"val": {"String": {"str": "r"}}, ` +
			`"location": 75}}, "result": {"A_Const": {"val": {"String": {"str": "table"}}, "location": 84}}, ` +
			`"location": 70}}, {"CaseWhen": {"expr": {"A_Const": {"val": {"String": {"str": "v"}}, ` +
			`"location": 97}}, "result": {"A_Const": {"val": {"String": {"str": "view"}}, "location": 106}}, ` +
			`"location": 92}}, {"CaseWhen": {"expr": {"A_Const": {"val": {"String": {"str": "m"}}, ` +
			`"location": 118}}, "result": {"A_Const": {"val": {"String": {"str": "materialized view"}}, ` +
			`"location": 127}}, "location": 113}}, {"CaseWhen": {"expr": {"A_Const": {"val": {"String": ` +
			`{"str": "i"}}, "location": 152}}, "result": {"A_Const": {"val": {"String": {"str": "index"}}, ` +
			`"location": 161}}, "location": 147}}, {"CaseWhen": {"expr": {"A_Const": {"val": {"String": ` +
			`{"str": "S"}}, "location": 174}}, "result": {"A_Const": {"val": {"String": {"str": "sequence"}}, ` +
			`"location": 183}}, "location": 169}}, {"CaseWhen": {"expr": {"A_Const": {"val": {"String": ` +
			`{"str": "s"}}, "location": 199}}, "result": {"A_Const": {"val": {"String": {"str": "special"}}, ` +
			`"location": 208}}, "location": 194}}, {"CaseWhen": {"expr": {"A_Const": {"val": {"String": ` +
			`{"str": "f"}}, "location": 223}}, "result": {"A_Const": {"val": {"String": {"str": "foreign table"}}, ` +
			`"location": 232}}, "location": 218}}], "location": 55}}, "location": 55}}, {"ResTarget": ` +
			`{"name": "Owner", "val": {"FuncCall": {"funcname": [{"String": {"str": "pg_catalog"}}, ` +
			`{"String": {"str": "pg_get_userbyid"}}], "args": [{"ColumnRef": {"fields": [{"String": ` +
			`{"str": "c"}}, {"String": {"str": "relowner"}}], "location": 292}}], "location": 265}}, "location": 265}}], ` +
			`"fromClause": [{"JoinExpr": {"jointype": 1, "larg": {"RangeVar": ` +
			`{"schemaname": "pg_catalog", "relname": "pg_class", "inhOpt": 2, "relpersistence": "p", ` +
			`"alias": {"Alias": {"aliasname": "c"}}, "location": 320}}, "rarg": {"RangeVar": {"schemaname": ` +
			`"pg_catalog", "relname": "pg_namespace", "inhOpt": 2, "relpersistence": "p", "alias": {"Alias": ` +
			`{"aliasname": "n"}}, "location": 357}}, "quals": {"A_Expr": {"kind": 0, "name": [{"String": ` +
			`{"str": "="}}], "lexpr": {"ColumnRef": {"fields": [{"String": {"str": "n"}}, {"String": ` +
			`{"str": "oid"}}], "location": 386}}, "rexpr": {"ColumnRef": {"fields": [{"String": {"str": "c"}}, ` +
			`{"String": {"str": "relnamespace"}}], "location": 394}}, "location": 392}}}}], ` +
			`"whereClause": {"A_Expr": {"kind": 1, "lexpr": {"A_Expr": {"kind": 1, "lexpr": {"A_Expr": ` +
			`{"kind": 1, "lexpr": {"A_Expr": {"kind": 1, "lexpr": {"A_Expr": {"kind": 9, "name": ` +
			`[{"String": {"str": "="}}], "lexpr": {"ColumnRef": {"fields": [{"String": {"str": "c"}}, ` +
			`{"String": {"str": "relkind"}}], "location": 415}}, "rexpr": [{"A_Const": {"val": {"String": ` +
			`{"str": "r"}}, "location": 429}}, {"A_Const": {"val": {"String": {"str": ""}}, "location": 433}}], ` +
			`"location": 425}}, "rexpr": {"A_Expr": {"kind": 0, "name": [{"String": {"str": "<>"}}], "lexpr": ` +
			`{"ColumnRef": {"fields": [{"String": {"str": "n"}}, {"String": {"str": "nspname"}}], "location": 447}}, ` +
			`"rexpr": {"A_Const": {"val": {"String": {"str": "pg_catalog"}}, "location": 460}}, "location": 457}}, ` +
			`"location": 443}}, "rexpr": {"A_Expr": {"kind": 0, "name": [{"String": {"str": "<>"}}], "lexpr": ` +
			`{"ColumnRef": {"fields": [{"String": {"str": "n"}}, {"String": {"str": "nspname"}}], "location": 483}}, ` +
			`"rexpr": {"A_Const": {"val": {"String": {"str": "information_schema"}}, "location": 496}}, "location": ` +
			`493}}, "location": 479}}, "rexpr": {"A_Expr": {"kind": 0, "name": [{"String": {"str": "!~"}}], "lexpr": ` +
			`{"ColumnRef": {"fields": [{"String": {"str": "n"}}, {"String": {"str": "nspname"}}], "location": 527}}, ` +
			`"rexpr": {"A_Const": {"val": {"String": {"str": "^pg_toast"}}, "location": 540}}, "location": 537}}, ` +
			`"location": 523}}, "rexpr": {"FuncCall": {"funcname": [{"String": {"str": "pg_catalog"}}, {"String": ` +
			`{"str": "pg_table_is_visible"}}], "args": [{"ColumnRef": {"fields": [{"String": {"str": "c"}}, ` +
			`{"String": {"str": "oid"}}], "location": 589}}], "location": 558}}, "location": 554}}, "sortClause": ` +
			`[{"SortBy": {"node": {"A_Const": {"val": {"Integer": {"ival": 1}}, "location": 605}}, "sortby_dir": 0, ` +
			`"sortby_nulls": 0, "location": -1}}, {"SortBy": {"node": {"A_Const": {"val": {"Integer": {"ival": 2}}, ` +
			`"location": 607}}, "sortby_dir": 0, "sortby_nulls": 0, "location": -1}}], "op": 0}}]`,
		pg_query.ParsetreeList{
			Statements: []nodes.Node{
				nodes.SelectStmt{
					TargetList: util.MakeListNode([]nodes.Node{
						nodes.ResTarget{
							Name: util.MakeStrPtr("Schema"),
							Val: nodes.ColumnRef{
								Fields: util.MakeListNode([]nodes.Node{
									util.MakeStrNode("n"),
									util.MakeStrNode("nspname"),
								}),
								Location: 7,
							},
							Location: 7,
						},
						nodes.ResTarget{
							Name: util.MakeStrPtr("Name"),
							Val: nodes.ColumnRef{
								Fields: util.MakeListNode([]nodes.Node{
									util.MakeStrNode("c"),
									util.MakeStrNode("relname"),
								}),
								Location: 32,
							},
							Location: 32,
						},
						nodes.ResTarget{
							Name: util.MakeStrPtr("Type"),
							Val: nodes.CaseExpr{
								Arg: nodes.ColumnRef{
									Fields: util.MakeListNode([]nodes.Node{
										util.MakeStrNode("c"),
										util.MakeStrNode("relkind"),
									}),
									Location: 60,
								},
								Args: util.MakeListNode([]nodes.Node{
									nodes.CaseWhen{
										Expr:     nodes.A_Const{Val: util.MakeStrNode("r"), Location: 75},
										Result:   nodes.A_Const{Val: util.MakeStrNode("table"), Location: 84},
										Location: 70,
									},
									nodes.CaseWhen{
										Expr:     nodes.A_Const{Val: util.MakeStrNode("v"), Location: 97},
										Result:   nodes.A_Const{Val: util.MakeStrNode("view"), Location: 106},
										Location: 92,
									},
									nodes.CaseWhen{
										Expr:     nodes.A_Const{Val: util.MakeStrNode("m"), Location: 118},
										Result:   nodes.A_Const{Val: util.MakeStrNode("materialized view"), Location: 127},
										Location: 113,
									},
									nodes.CaseWhen{
										Expr:     nodes.A_Const{Val: util.MakeStrNode("i"), Location: 152},
										Result:   nodes.A_Const{Val: util.MakeStrNode("index"), Location: 161},
										Location: 147,
									},
									nodes.CaseWhen{
										Expr:     nodes.A_Const{Val: util.MakeStrNode("S"), Location: 174},
										Result:   nodes.A_Const{Val: util.MakeStrNode("sequence"), Location: 183},
										Location: 169,
									},
									nodes.CaseWhen{
										Expr:     nodes.A_Const{Val: util.MakeStrNode("s"), Location: 199},
										Result:   nodes.A_Const{Val: util.MakeStrNode("special"), Location: 208},
										Location: 194,
									},
									nodes.CaseWhen{
										Expr:     nodes.A_Const{Val: util.MakeStrNode("f"), Location: 223},
										Result:   nodes.A_Const{Val: util.MakeStrNode("foreign table"), Location: 232},
										Location: 218,
									},
								}),
								Location: 55,
							},
							Location: 55,
						},
						nodes.ResTarget{
							Name: util.MakeStrPtr("Owner"),
							Val: nodes.FuncCall{
								Funcname: util.MakeListNode([]nodes.Node{
									util.MakeStrNode("pg_catalog"),
									util.MakeStrNode("pg_get_userbyid"),
								}),
								Args: util.MakeListNode([]nodes.Node{
									nodes.ColumnRef{
										Fields: util.MakeListNode([]nodes.Node{
											util.MakeStrNode("c"),
											util.MakeStrNode("relowner"),
										}),
										Location: 292,
									},
								}),
								Location: 265,
							},
							Location: 265,
						},
					}),
					FromClause: util.MakeListNode([]nodes.Node{
						nodes.JoinExpr{
							Jointype: 1,
							Larg: nodes.RangeVar{
								Schemaname:     util.MakeStrPtr("pg_catalog"),
								Relname:        util.MakeStrPtr("pg_class"),
								InhOpt:         nodes.INH_DEFAULT,
								Relpersistence: 'p',
								Alias: &nodes.Alias{
									Aliasname: util.MakeStrPtr("c"),
								},
								Location: 320,
							},
							Rarg: nodes.RangeVar{
								Schemaname:     util.MakeStrPtr("pg_catalog"),
								Relname:        util.MakeStrPtr("pg_namespace"),
								InhOpt:         nodes.INH_DEFAULT,
								Relpersistence: 'p',
								Alias: &nodes.Alias{
									Aliasname: util.MakeStrPtr("n"),
								},
								Location: 357,
							},
							Quals: nodes.A_Expr{
								Kind: nodes.AEXPR_OP,
								Name: util.MakeListNode([]nodes.Node{util.MakeStrNode("=")}),
								Lexpr: nodes.ColumnRef{
									Fields:   util.MakeListNode([]nodes.Node{util.MakeStrNode("n"), util.MakeStrNode("oid")}),
									Location: 386,
								},
								Rexpr: nodes.ColumnRef{
									Fields:   util.MakeListNode([]nodes.Node{util.MakeStrNode("c"), util.MakeStrNode("relnamespace")}),
									Location: 394,
								},
								Location: 392,
							},
						},
					}),
					WhereClause: nodes.A_Expr{
						Kind: nodes.AEXPR_AND,
						Lexpr: nodes.A_Expr{
							Kind: nodes.AEXPR_AND,
							Lexpr: nodes.A_Expr{
								Kind: nodes.AEXPR_AND,
								Lexpr: nodes.A_Expr{
									Kind: nodes.AEXPR_AND,
									Lexpr: nodes.A_Expr{
										Kind: nodes.AEXPR_IN,
										Name: util.MakeListNode([]nodes.Node{util.MakeStrNode("=")}),
										Lexpr: nodes.ColumnRef{
											Fields:   util.MakeListNode([]nodes.Node{util.MakeStrNode("c"), util.MakeStrNode("relkind")}),
											Location: 415,
										},
										Rexpr: util.MakeListNode([]nodes.Node{
											nodes.A_Const{Val: util.MakeStrNode("r"), Location: 429},
											nodes.A_Const{Val: util.MakeStrNode(""), Location: 433},
										}),
										Location: 425,
									},
									Rexpr: nodes.A_Expr{
										Name: util.MakeListNode([]nodes.Node{util.MakeStrNode("<>")}),
										Lexpr: nodes.ColumnRef{
											Fields:   util.MakeListNode([]nodes.Node{util.MakeStrNode("n"), util.MakeStrNode("nspname")}),
											Location: 447,
										},
										Rexpr:    nodes.A_Const{Val: util.MakeStrNode("pg_catalog"), Location: 460},
										Location: 457,
									},
									Location: 443,
								},
								Rexpr: nodes.A_Expr{
									Kind: nodes.AEXPR_OP,
									Name: util.MakeListNode([]nodes.Node{util.MakeStrNode("<>")}),
									Lexpr: nodes.ColumnRef{
										Fields:   util.MakeListNode([]nodes.Node{util.MakeStrNode("n"), util.MakeStrNode("nspname")}),
										Location: 483,
									},
									Rexpr:    nodes.A_Const{Val: util.MakeStrNode("information_schema"), Location: 496},
									Location: 493,
								},
								Location: 479,
							},
							Rexpr: nodes.A_Expr{
								Kind: nodes.AEXPR_OP,
								Name: util.MakeListNode([]nodes.Node{util.MakeStrNode("!~")}),
								Lexpr: nodes.ColumnRef{
									Fields:   util.MakeListNode([]nodes.Node{util.MakeStrNode("n"), util.MakeStrNode("nspname")}),
									Location: 527,
								},
								Rexpr: nodes.A_Const{
									Val:      util.MakeStrNode("^pg_toast"),
									Location: 540,
								},
								Location: 537,
							},
							Location: 523,
						},
						Rexpr: nodes.FuncCall{
							Funcname: util.MakeListNode([]nodes.Node{
								util.MakeStrNode("pg_catalog"),
								util.MakeStrNode("pg_table_is_visible"),
							}),
							Args: util.MakeListNode([]nodes.Node{
								nodes.ColumnRef{
									Fields:   util.MakeListNode([]nodes.Node{util.MakeStrNode("c"), util.MakeStrNode("oid")}),
									Location: 589,
								},
							}),
							Location: 558,
						},
						Location: 554,
					},
					SortClause: util.MakeListNode([]nodes.Node{
						nodes.SortBy{
							Node: nodes.A_Const{
								Val:      util.MakeIntNode(1),
								Location: 605,
							},
							Location: -1,
						},
						nodes.SortBy{
							Node: nodes.A_Const{
								Val:      util.MakeIntNode(2),
								Location: 607,
							},
							Location: -1,
						},
					}),
				},
			},
		},
	},
	{
		`CREATE FUNCTION change_trigger_v2() RETURNS trigger
  LANGUAGE plpgsql
  AS $$
    DECLARE
    BEGIN
      PERFORM 'dummy';
    END;
    $$;`,
		`[{"CreateFunctionStmt": {"funcname": [{"String": {"str": ` +
			`"change_trigger_v2"}}], "returnType": {"TypeName": {"names": [{"String": ` +
			`{"str": "trigger"}}], "typemod": -1, "location": 44}}, "options": [{"DefElem": {"defname": "language", ` +
			`"arg": {"String": {"str": "plpgsql"}}, "defaction": 0}}, {"DefElem": {"defname": ` +
			`"as", "arg": [{"String": {"str": "\n    DECLARE\n    BEGIN\n      PERFORM 'dummy';\n    END;\n    "}}], ` +
			`"defaction": 0}}]}}]`,
		pg_query.ParsetreeList{
			Statements: []nodes.Node{
				nodes.CreateFunctionStmt{
					Funcname: util.MakeListNode([]nodes.Node{
						util.MakeStrNode("change_trigger_v2"),
					}),
					ReturnType: &nodes.TypeName{
						Names: util.MakeListNode([]nodes.Node{
							util.MakeStrNode("trigger"),
						}),
						Typemod:  -1,
						Location: 44,
					},
					Options: util.MakeListNode([]nodes.Node{
						nodes.DefElem{
							Defname: util.MakeStrPtr("language"),
							Arg:     util.MakeStrNode("plpgsql"),
						},
						nodes.DefElem{
							Defname: util.MakeStrPtr("as"),
							Arg:     util.MakeListNode([]nodes.Node{util.MakeStrNode("\n    DECLARE\n    BEGIN\n      PERFORM 'dummy';\n    END;\n    ")}),
						},
					}),
				},
			},
		},
	},
	{
		`CREATE TABLE test (
    id SERIAL PRIMARY KEY,
    user_id integer DEFAULT 0 NOT NULL,
    created_at timestamp without time zone NOT NULL);`,
		`[{"CreateStmt": {"relation": {"RangeVar": {"relname": "test", "inhOpt": 2, ` +
			`"relpersistence": "p", "location": 13}}, "tableElts": [{"ColumnDef": {"colname": ` +
			`"id", "typeName": {"TypeName": {"names": [{"String": {"str": "serial"}}], ` +
			`"typemod": -1, "location": 27}}, "is_local": true, "constraints": [{"Constraint": ` +
			`{"contype": 4, "location": 34}}], "location": 24}}, {"ColumnDef": {"colname": ` +
			`"user_id", "typeName": {"TypeName": {"names": [{"String": {"str": "pg_catalog"}}, ` +
			`{"String": {"str": "int4"}}], "typemod": -1, "location": 59}}, "is_local": true, ` +
			`"constraints": [{"Constraint": {"contype": 2, "location": 67, "raw_expr": {"A_Const": ` +
			`{"val": {"Integer": {"ival": 0}}, "location": 75}}}}, {"Constraint": {"contype": 1, ` +
			`"location": 77}}], "location": 51}}, {"ColumnDef": {"colname": "created_at", "typeName": ` +
			`{"TypeName": {"names": [{"String": {"str": "pg_catalog"}}, {"String": {"str": ` +
			`"timestamp"}}], "typemod": -1, "location": 102}}, "is_local": true, "constraints": ` +
			`[{"Constraint": {"contype": 1, "location": 130}}], "location": 91}}], "oncommit": 0}}]`,
		pg_query.ParsetreeList{
			Statements: []nodes.Node{
				nodes.CreateStmt{
					Relation: &nodes.RangeVar{
						Relname:        util.MakeStrPtr("test"),
						InhOpt:         nodes.INH_DEFAULT,
						Relpersistence: 'p',
						Location:       13,
					},
					TableElts: util.MakeListNode([]nodes.Node{
						nodes.ColumnDef{
							Colname: util.MakeStrPtr("id"),
							TypeName: &nodes.TypeName{
								Names: util.MakeListNode([]nodes.Node{
									util.MakeStrNode("serial"),
								}),
								Typemod:  -1,
								Location: 27,
							},
							IsLocal: true,
							Constraints: util.MakeListNode([]nodes.Node{
								nodes.Constraint{
									Contype:  nodes.CONSTR_PRIMARY,
									Location: 34,
								},
							}),
							Location: 24,
						},
						nodes.ColumnDef{
							Colname: util.MakeStrPtr("user_id"),
							TypeName: &nodes.TypeName{
								Names: util.MakeListNode([]nodes.Node{
									util.MakeStrNode("pg_catalog"),
									util.MakeStrNode("int4"),
								}),
								Typemod:  -1,
								Location: 59,
							},
							IsLocal: true,
							Constraints: util.MakeListNode([]nodes.Node{
								nodes.Constraint{
									Contype:  nodes.CONSTR_DEFAULT,
									Location: 67,
									RawExpr: nodes.A_Const{
										Val:      util.MakeIntNode(0),
										Location: 75,
									},
								},
								nodes.Constraint{
									Contype:  nodes.CONSTR_NOTNULL,
									Location: 77,
								},
							}),
							Location: 51,
						},
						nodes.ColumnDef{
							Colname: util.MakeStrPtr("created_at"),
							TypeName: &nodes.TypeName{
								Names: util.MakeListNode([]nodes.Node{
									util.MakeStrNode("pg_catalog"),
									util.MakeStrNode("timestamp"),
								}),
								Typemod:  -1,
								Location: 102,
							},
							IsLocal: true,
							Constraints: util.MakeListNode([]nodes.Node{
								nodes.Constraint{
									Contype:  nodes.CONSTR_NOTNULL,
									Location: 130,
								},
							}),
							Location: 91,
						},
					}),
					Oncommit: nodes.ONCOMMIT_NOOP,
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
