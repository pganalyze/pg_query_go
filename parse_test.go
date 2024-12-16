//go:build cgo
// +build cgo

package pg_query_test

import (
	"fmt"
	"reflect"
	"sync"
	"testing"

	"github.com/google/go-cmp/cmp"
	pg_query "github.com/xataio/pg_query_go/v6"
	"github.com/xataio/pg_query_go/v6/parser"
	"google.golang.org/protobuf/testing/protocmp"
)

var parseTests = []struct {
	input        string
	expectedJSON string
	expectedTree *pg_query.ParseResult
}{
	{
		"SELECT 1",
		`{"version":170004,"stmts":[{"stmt":{"SelectStmt":{"targetList":[{"ResTarget":{"val":{"A_Const":{"ival":{"ival":1},"location":7}},"location":7}}],"limitOption":"LIMIT_OPTION_DEFAULT","op":"SETOP_NONE"}}}]}`,
		&pg_query.ParseResult{
			Version: int32(170004),
			Stmts: []*pg_query.RawStmt{
				{
					Stmt: &pg_query.Node{
						Node: &pg_query.Node_SelectStmt{
							SelectStmt: &pg_query.SelectStmt{
								LimitOption: pg_query.LimitOption_LIMIT_OPTION_DEFAULT,
								Op:          pg_query.SetOperation_SETOP_NONE,
								TargetList: []*pg_query.Node{
									pg_query.MakeResTargetNodeWithVal(
										pg_query.MakeAConstIntNode(1, 7),
										7,
									),
								},
							},
						},
					},
				},
			},
		},
	},
	{
		"SELECT * FROM x WHERE z = 1",
		`{"version":170004,"stmts":[{"stmt":{"SelectStmt":{"targetList":[{"ResTarget":{"val":{"ColumnRef":{"fields":[{"A_Star":{}}],"location":7}},"location":7}}],"fromClause":[{"RangeVar":{"relname":"x","inh":true,"relpersistence":"p","location":14}}],"whereClause":{"A_Expr":{"kind":"AEXPR_OP","name":[{"String":{"sval":"="}}],"lexpr":{"ColumnRef":{"fields":[{"String":{"sval":"z"}}],"location":22}},"rexpr":{"A_Const":{"ival":{"ival":1},"location":26}},"location":24}},"limitOption":"LIMIT_OPTION_DEFAULT","op":"SETOP_NONE"}}}]}`,
		&pg_query.ParseResult{
			Version: int32(170004),
			Stmts: []*pg_query.RawStmt{
				{
					Stmt: &pg_query.Node{
						Node: &pg_query.Node_SelectStmt{
							SelectStmt: &pg_query.SelectStmt{
								LimitOption: pg_query.LimitOption_LIMIT_OPTION_DEFAULT,
								Op:          pg_query.SetOperation_SETOP_NONE,
								TargetList: []*pg_query.Node{
									pg_query.MakeResTargetNodeWithVal(
										pg_query.MakeColumnRefNode(
											[]*pg_query.Node{
												pg_query.MakeAStarNode(),
											},
											7,
										),
										7,
									),
								},
								FromClause: []*pg_query.Node{
									pg_query.MakeSimpleRangeVarNode("x", 14),
								},
								WhereClause: &pg_query.Node{
									Node: &pg_query.Node_AExpr{
										AExpr: &pg_query.A_Expr{
											Kind: pg_query.A_Expr_Kind_AEXPR_OP,
											Name: []*pg_query.Node{
												pg_query.MakeStrNode("="),
											},
											Lexpr: pg_query.MakeColumnRefNode(
												[]*pg_query.Node{
													pg_query.MakeStrNode("z"),
												},
												22,
											),
											Rexpr:    pg_query.MakeAConstIntNode(1, 26),
											Location: 24,
										},
									},
								},
							},
						},
					},
				},
			},
		},
	},
	{
		`INSERT INTO "schema_index_stats" ("snapshot_id","schema_index_id","size_bytes") VALUES (11710849,8448632,16384),(11710849,8448633,16384) RETURNING id`,
		`{"version":170004,"stmts":[{"stmt":{"InsertStmt":{"relation":{"relname":"schema_index_stats","inh":true,"relpersistence":"p","location":12},"cols":[{"ResTarget":{"name":"snapshot_id","location":34}},{"ResTarget":{"name":"schema_index_id","location":48}},{"ResTarget":{"name":"size_bytes","location":66}}],"selectStmt":{"SelectStmt":{"valuesLists":[{"List":{"items":[{"A_Const":{"ival":{"ival":11710849},"location":88}},{"A_Const":{"ival":{"ival":8448632},"location":97}},{"A_Const":{"ival":{"ival":16384},"location":105}}]}},{"List":{"items":[{"A_Const":{"ival":{"ival":11710849},"location":113}},{"A_Const":{"ival":{"ival":8448633},"location":122}},{"A_Const":{"ival":{"ival":16384},"location":130}}]}}],"limitOption":"LIMIT_OPTION_DEFAULT","op":"SETOP_NONE"}},"returningList":[{"ResTarget":{"val":{"ColumnRef":{"fields":[{"String":{"sval":"id"}}],"location":147}},"location":147}}],"override":"OVERRIDING_NOT_SET"}}}]}`,
		&pg_query.ParseResult{
			Version: int32(170004),
			Stmts: []*pg_query.RawStmt{
				{
					Stmt: &pg_query.Node{
						Node: &pg_query.Node_InsertStmt{
							InsertStmt: &pg_query.InsertStmt{
								Relation: pg_query.MakeSimpleRangeVar("schema_index_stats", 12),
								Cols: []*pg_query.Node{
									pg_query.MakeResTargetNodeWithName("snapshot_id", 34),
									pg_query.MakeResTargetNodeWithName("schema_index_id", 48),
									pg_query.MakeResTargetNodeWithName("size_bytes", 66),
								},
								Override: pg_query.OverridingKind_OVERRIDING_NOT_SET,
								SelectStmt: &pg_query.Node{
									Node: &pg_query.Node_SelectStmt{
										SelectStmt: &pg_query.SelectStmt{
											LimitOption: pg_query.LimitOption_LIMIT_OPTION_DEFAULT,
											Op:          pg_query.SetOperation_SETOP_NONE,
											ValuesLists: []*pg_query.Node{
												pg_query.MakeListNode([]*pg_query.Node{
													pg_query.MakeAConstIntNode(11710849, 88),
													pg_query.MakeAConstIntNode(8448632, 97),
													pg_query.MakeAConstIntNode(16384, 105),
												}),
												pg_query.MakeListNode([]*pg_query.Node{
													pg_query.MakeAConstIntNode(11710849, 113),
													pg_query.MakeAConstIntNode(8448633, 122),
													pg_query.MakeAConstIntNode(16384, 130),
												}),
											},
										},
									},
								},
								ReturningList: []*pg_query.Node{
									pg_query.MakeResTargetNodeWithVal(
										pg_query.MakeColumnRefNode(
											[]*pg_query.Node{
												pg_query.MakeStrNode("id"),
											},
											147,
										),
										147,
									),
								},
							},
						},
					},
				},
			},
		},
	},
	{
		"SELECT * FROM x WHERE y IN ($1)",
		`{"version":170004,"stmts":[{"stmt":{"SelectStmt":{"targetList":[{"ResTarget":{"val":{"ColumnRef":{"fields":[{"A_Star":{}}],"location":7}},"location":7}}],"fromClause":[{"RangeVar":{"relname":"x","inh":true,"relpersistence":"p","location":14}}],"whereClause":{"A_Expr":{"kind":"AEXPR_IN","name":[{"String":{"sval":"="}}],"lexpr":{"ColumnRef":{"fields":[{"String":{"sval":"y"}}],"location":22}},"rexpr":{"List":{"items":[{"ParamRef":{"number":1,"location":28}}]}},"location":24}},"limitOption":"LIMIT_OPTION_DEFAULT","op":"SETOP_NONE"}}}]}`,
		&pg_query.ParseResult{
			Version: int32(170004),
			Stmts: []*pg_query.RawStmt{
				{
					Stmt: &pg_query.Node{
						Node: &pg_query.Node_SelectStmt{
							SelectStmt: &pg_query.SelectStmt{
								LimitOption: pg_query.LimitOption_LIMIT_OPTION_DEFAULT,
								Op:          pg_query.SetOperation_SETOP_NONE,
								TargetList: []*pg_query.Node{
									pg_query.MakeResTargetNodeWithVal(
										pg_query.MakeColumnRefNode(
											[]*pg_query.Node{
												pg_query.MakeAStarNode(),
											},
											7,
										),
										7,
									),
								},
								FromClause: []*pg_query.Node{
									pg_query.MakeSimpleRangeVarNode("x", 14),
								},
								WhereClause: &pg_query.Node{
									Node: &pg_query.Node_AExpr{
										AExpr: &pg_query.A_Expr{
											Kind: pg_query.A_Expr_Kind_AEXPR_IN,
											Name: []*pg_query.Node{
												pg_query.MakeStrNode("="),
											},
											Lexpr: pg_query.MakeColumnRefNode(
												[]*pg_query.Node{
													pg_query.MakeStrNode("y"),
												},
												22,
											),
											Rexpr: pg_query.MakeListNode([]*pg_query.Node{
												pg_query.MakeParamRefNode(1, 28),
											}),
											Location: 24,
										},
									},
								},
							},
						},
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
		`{"version":170004,"stmts":[{"stmt":{"SelectStmt":{"targetList":[{"ResTarget":{"name":"Schema","val":{"ColumnRef":{"fields":[{"String":{"sval":"n"}},{"String":{"sval":"nspname"}}],"location":7}},"location":7}},{"ResTarget":{"name":"Name","val":{"ColumnRef":{"fields":[{"String":{"sval":"c"}},{"String":{"sval":"relname"}}],"location":36}},"location":36}},{"ResTarget":{"name":"Type","val":{"CaseExpr":{"arg":{"ColumnRef":{"fields":[{"String":{"sval":"c"}},{"String":{"sval":"relkind"}}],"location":68}},"args":[{"CaseWhen":{"expr":{"A_Const":{"sval":{"sval":"r"},"location":83}},"result":{"A_Const":{"sval":{"sval":"table"},"location":92}},"location":78}},{"CaseWhen":{"expr":{"A_Const":{"sval":{"sval":"v"},"location":105}},"result":{"A_Const":{"sval":{"sval":"view"},"location":114}},"location":100}},{"CaseWhen":{"expr":{"A_Const":{"sval":{"sval":"m"},"location":126}},"result":{"A_Const":{"sval":{"sval":"materialized view"},"location":135}},"location":121}},{"CaseWhen":{"expr":{"A_Const":{"sval":{"sval":"i"},"location":160}},"result":{"A_Const":{"sval":{"sval":"index"},"location":169}},"location":155}},{"CaseWhen":{"expr":{"A_Const":{"sval":{"sval":"S"},"location":182}},"result":{"A_Const":{"sval":{"sval":"sequence"},"location":191}},"location":177}},{"CaseWhen":{"expr":{"A_Const":{"sval":{"sval":"s"},"location":207}},"result":{"A_Const":{"sval":{"sval":"special"},"location":216}},"location":202}},{"CaseWhen":{"expr":{"A_Const":{"sval":{"sval":"f"},"location":231}},"result":{"A_Const":{"sval":{"sval":"foreign table"},"location":240}},"location":226}}],"location":63}},"location":63}},{"ResTarget":{"name":"Owner","val":{"FuncCall":{"funcname":[{"String":{"sval":"pg_catalog"}},{"String":{"sval":"pg_get_userbyid"}}],"args":[{"ColumnRef":{"fields":[{"String":{"sval":"c"}},{"String":{"sval":"relowner"}}],"location":304}}],"funcformat":"COERCE_EXPLICIT_CALL","location":277}},"location":277}}],"fromClause":[{"JoinExpr":{"jointype":"JOIN_LEFT","larg":{"RangeVar":{"schemaname":"pg_catalog","relname":"pg_class","inh":true,"relpersistence":"p","alias":{"aliasname":"c"},"location":336}},"rarg":{"RangeVar":{"schemaname":"pg_catalog","relname":"pg_namespace","inh":true,"relpersistence":"p","alias":{"aliasname":"n"},"location":374}},"quals":{"A_Expr":{"kind":"AEXPR_OP","name":[{"String":{"sval":"="}}],"lexpr":{"ColumnRef":{"fields":[{"String":{"sval":"n"}},{"String":{"sval":"oid"}}],"location":403}},"rexpr":{"ColumnRef":{"fields":[{"String":{"sval":"c"}},{"String":{"sval":"relnamespace"}}],"location":411}},"location":409}}}}],"whereClause":{"BoolExpr":{"boolop":"AND_EXPR","args":[{"A_Expr":{"kind":"AEXPR_IN","name":[{"String":{"sval":"="}}],"lexpr":{"ColumnRef":{"fields":[{"String":{"sval":"c"}},{"String":{"sval":"relkind"}}],"location":435}},"rexpr":{"List":{"items":[{"A_Const":{"sval":{"sval":"r"},"location":449}},{"A_Const":{"sval":{"sval":""},"location":453}}]}},"location":445}},{"A_Expr":{"kind":"AEXPR_OP","name":[{"String":{"sval":"\u003c\u003e"}}],"lexpr":{"ColumnRef":{"fields":[{"String":{"sval":"n"}},{"String":{"sval":"nspname"}}],"location":467}},"rexpr":{"A_Const":{"sval":{"sval":"pg_catalog"},"location":480}},"location":477}},{"A_Expr":{"kind":"AEXPR_OP","name":[{"String":{"sval":"\u003c\u003e"}}],"lexpr":{"ColumnRef":{"fields":[{"String":{"sval":"n"}},{"String":{"sval":"nspname"}}],"location":503}},"rexpr":{"A_Const":{"sval":{"sval":"information_schema"},"location":516}},"location":513}},{"A_Expr":{"kind":"AEXPR_OP","name":[{"String":{"sval":"!~"}}],"lexpr":{"ColumnRef":{"fields":[{"String":{"sval":"n"}},{"String":{"sval":"nspname"}}],"location":547}},"rexpr":{"A_Const":{"sval":{"sval":"^pg_toast"},"location":560}},"location":557}},{"FuncCall":{"funcname":[{"String":{"sval":"pg_catalog"}},{"String":{"sval":"pg_table_is_visible"}}],"args":[{"ColumnRef":{"fields":[{"String":{"sval":"c"}},{"String":{"sval":"oid"}}],"location":613}}],"funcformat":"COERCE_EXPLICIT_CALL","location":582}}],"location":463}},"sortClause":[{"SortBy":{"node":{"A_Const":{"ival":{"ival":1},"location":632}},"sortby_dir":"SORTBY_DEFAULT","sortby_nulls":"SORTBY_NULLS_DEFAULT","location":-1}},{"SortBy":{"node":{"A_Const":{"ival":{"ival":2},"location":634}},"sortby_dir":"SORTBY_DEFAULT","sortby_nulls":"SORTBY_NULLS_DEFAULT","location":-1}}],"limitOption":"LIMIT_OPTION_DEFAULT","op":"SETOP_NONE"}},"stmt_len":635}]}`,
		&pg_query.ParseResult{
			Version: int32(170004),
			Stmts: []*pg_query.RawStmt{
				{
					Stmt: &pg_query.Node{
						Node: &pg_query.Node_SelectStmt{
							SelectStmt: &pg_query.SelectStmt{
								LimitOption: pg_query.LimitOption_LIMIT_OPTION_DEFAULT,
								Op:          pg_query.SetOperation_SETOP_NONE,
								TargetList: []*pg_query.Node{
									pg_query.MakeResTargetNodeWithNameAndVal(
										"Schema",
										pg_query.MakeColumnRefNode(
											[]*pg_query.Node{
												pg_query.MakeStrNode("n"),
												pg_query.MakeStrNode("nspname"),
											},
											7,
										),
										7,
									),
									pg_query.MakeResTargetNodeWithNameAndVal(
										"Name",
										pg_query.MakeColumnRefNode(
											[]*pg_query.Node{
												pg_query.MakeStrNode("c"),
												pg_query.MakeStrNode("relname"),
											},
											36,
										),
										36,
									),
									pg_query.MakeResTargetNodeWithNameAndVal(
										"Type",
										pg_query.MakeCaseExprNode(
											pg_query.MakeColumnRefNode(
												[]*pg_query.Node{
													pg_query.MakeStrNode("c"),
													pg_query.MakeStrNode("relkind"),
												},
												68,
											),
											[]*pg_query.Node{
												pg_query.MakeCaseWhenNode(pg_query.MakeAConstStrNode("r", 83), pg_query.MakeAConstStrNode("table", 92), 78),
												pg_query.MakeCaseWhenNode(pg_query.MakeAConstStrNode("v", 105), pg_query.MakeAConstStrNode("view", 114), 100),
												pg_query.MakeCaseWhenNode(pg_query.MakeAConstStrNode("m", 126), pg_query.MakeAConstStrNode("materialized view", 135), 121),
												pg_query.MakeCaseWhenNode(pg_query.MakeAConstStrNode("i", 160), pg_query.MakeAConstStrNode("index", 169), 155),
												pg_query.MakeCaseWhenNode(pg_query.MakeAConstStrNode("S", 182), pg_query.MakeAConstStrNode("sequence", 191), 177),
												pg_query.MakeCaseWhenNode(pg_query.MakeAConstStrNode("s", 207), pg_query.MakeAConstStrNode("special", 216), 202),
												pg_query.MakeCaseWhenNode(pg_query.MakeAConstStrNode("f", 231), pg_query.MakeAConstStrNode("foreign table", 240), 226),
											},
											63,
										),
										63,
									),
									pg_query.MakeResTargetNodeWithNameAndVal(
										"Owner",
										pg_query.MakeFuncCallNode(
											[]*pg_query.Node{
												pg_query.MakeStrNode("pg_catalog"),
												pg_query.MakeStrNode("pg_get_userbyid"),
											},
											[]*pg_query.Node{
												pg_query.MakeColumnRefNode(
													[]*pg_query.Node{
														pg_query.MakeStrNode("c"),
														pg_query.MakeStrNode("relowner"),
													},
													304,
												),
											},
											277,
										),
										277,
									),
								},
								FromClause: []*pg_query.Node{
									pg_query.MakeJoinExprNode(
										pg_query.JoinType_JOIN_LEFT,
										pg_query.MakeFullRangeVarNode("pg_catalog", "pg_class", "c", 336),
										pg_query.MakeFullRangeVarNode("pg_catalog", "pg_namespace", "n", 374),
										pg_query.MakeAExprNode(
											pg_query.A_Expr_Kind_AEXPR_OP,
											[]*pg_query.Node{
												pg_query.MakeStrNode("="),
											},
											pg_query.MakeColumnRefNode(
												[]*pg_query.Node{pg_query.MakeStrNode("n"), pg_query.MakeStrNode("oid")},
												403,
											),
											pg_query.MakeColumnRefNode(
												[]*pg_query.Node{pg_query.MakeStrNode("c"), pg_query.MakeStrNode("relnamespace")},
												411,
											),
											409,
										),
									),
								},
								WhereClause: pg_query.MakeBoolExprNode(
									pg_query.BoolExprType_AND_EXPR,
									[]*pg_query.Node{
										pg_query.MakeAExprNode(
											pg_query.A_Expr_Kind_AEXPR_IN,
											[]*pg_query.Node{pg_query.MakeStrNode("=")},
											pg_query.MakeColumnRefNode(
												[]*pg_query.Node{pg_query.MakeStrNode("c"), pg_query.MakeStrNode("relkind")},
												435,
											),
											pg_query.MakeListNode([]*pg_query.Node{
												pg_query.MakeAConstStrNode("r", 449),
												pg_query.MakeAConstStrNode("", 453),
											}),
											445,
										),
										pg_query.MakeAExprNode(
											pg_query.A_Expr_Kind_AEXPR_OP,
											[]*pg_query.Node{pg_query.MakeStrNode("<>")},
											pg_query.MakeColumnRefNode(
												[]*pg_query.Node{pg_query.MakeStrNode("n"), pg_query.MakeStrNode("nspname")},
												467,
											),
											pg_query.MakeAConstStrNode("pg_catalog", 480),
											477,
										),
										pg_query.MakeAExprNode(
											pg_query.A_Expr_Kind_AEXPR_OP,
											[]*pg_query.Node{pg_query.MakeStrNode("<>")},
											pg_query.MakeColumnRefNode(
												[]*pg_query.Node{pg_query.MakeStrNode("n"), pg_query.MakeStrNode("nspname")},
												503,
											),
											pg_query.MakeAConstStrNode("information_schema", 516),
											513,
										),
										pg_query.MakeAExprNode(
											pg_query.A_Expr_Kind_AEXPR_OP,
											[]*pg_query.Node{pg_query.MakeStrNode("!~")},
											pg_query.MakeColumnRefNode(
												[]*pg_query.Node{pg_query.MakeStrNode("n"), pg_query.MakeStrNode("nspname")},
												547,
											),
											pg_query.MakeAConstStrNode("^pg_toast", 560),
											557,
										),
										pg_query.MakeFuncCallNode(
											[]*pg_query.Node{
												pg_query.MakeStrNode("pg_catalog"),
												pg_query.MakeStrNode("pg_table_is_visible"),
											},
											[]*pg_query.Node{
												pg_query.MakeColumnRefNode(
													[]*pg_query.Node{pg_query.MakeStrNode("c"), pg_query.MakeStrNode("oid")},
													613,
												),
											},
											582,
										),
									},
									463,
								),
								SortClause: []*pg_query.Node{
									pg_query.MakeSortByNode(pg_query.MakeAConstIntNode(1, 632), pg_query.SortByDir_SORTBY_DEFAULT, pg_query.SortByNulls_SORTBY_NULLS_DEFAULT, -1),
									pg_query.MakeSortByNode(pg_query.MakeAConstIntNode(2, 634), pg_query.SortByDir_SORTBY_DEFAULT, pg_query.SortByNulls_SORTBY_NULLS_DEFAULT, -1),
								},
							},
						},
					},
					StmtLen: 635,
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
		`{"version":170004,"stmts":[{"stmt":{"CreateFunctionStmt":{"funcname":[{"String":{"sval":"change_trigger_v2"}}],"returnType":{"names":[{"String":{"sval":"trigger"}}],"typemod":-1,"location":44},"options":[{"DefElem":{"defname":"language","arg":{"String":{"sval":"plpgsql"}},"defaction":"DEFELEM_UNSPEC","location":53}},{"DefElem":{"defname":"as","arg":{"List":{"items":[{"String":{"sval":"\n\t\tDECLARE\n\t\tBEGIN\n\t\t\tPERFORM 'dummy';\n\t\tEND;\n\t\t"}}]}},"defaction":"DEFELEM_UNSPEC","location":71}}]}},"stmt_len":126}]}`,
		&pg_query.ParseResult{
			Version: int32(170004),
			Stmts: []*pg_query.RawStmt{
				{
					Stmt: &pg_query.Node{
						Node: &pg_query.Node_CreateFunctionStmt{
							CreateFunctionStmt: &pg_query.CreateFunctionStmt{
								Funcname: []*pg_query.Node{
									pg_query.MakeStrNode("change_trigger_v2"),
								},
								ReturnType: &pg_query.TypeName{
									Names: []*pg_query.Node{
										pg_query.MakeStrNode("trigger"),
									},
									Typemod:  -1,
									Location: 44,
								},
								Options: []*pg_query.Node{
									pg_query.MakeSimpleDefElemNode("language", pg_query.MakeStrNode("plpgsql"), 53),
									pg_query.MakeSimpleDefElemNode(
										"as",
										pg_query.MakeListNode(
											[]*pg_query.Node{
												pg_query.MakeStrNode("\n\t\tDECLARE\n\t\tBEGIN\n\t\t\tPERFORM 'dummy';\n\t\tEND;\n\t\t"),
											},
										),
										71,
									),
								},
							},
						},
					},
					StmtLen: 126,
				},
			},
		},
	},
	{
		`CREATE TABLE test (
			 id SERIAL PRIMARY KEY,
			 user_id integer DEFAULT 0 NOT NULL,
			 created_at timestamp without time zone NOT NULL);`,
		`{"version":170004,"stmts":[{"stmt":{"CreateStmt":{"relation":{"relname":"test","inh":true,"relpersistence":"p","location":13},"tableElts":[{"ColumnDef":{"colname":"id","typeName":{"names":[{"String":{"sval":"serial"}}],"typemod":-1,"location":27},"is_local":true,"constraints":[{"Constraint":{"contype":"CONSTR_PRIMARY","location":34}}],"location":24}},{"ColumnDef":{"colname":"user_id","typeName":{"names":[{"String":{"sval":"pg_catalog"}},{"String":{"sval":"int4"}}],"typemod":-1,"location":59},"is_local":true,"constraints":[{"Constraint":{"contype":"CONSTR_DEFAULT","raw_expr":{"A_Const":{"ival":{},"location":75}},"location":67}},{"Constraint":{"contype":"CONSTR_NOTNULL","location":77}}],"location":51}},{"ColumnDef":{"colname":"created_at","typeName":{"names":[{"String":{"sval":"pg_catalog"}},{"String":{"sval":"timestamp"}}],"typemod":-1,"location":102},"is_local":true,"constraints":[{"Constraint":{"contype":"CONSTR_NOTNULL","location":130}}],"location":91}}],"oncommit":"ONCOMMIT_NOOP"}},"stmt_len":139}]}`,
		&pg_query.ParseResult{
			Version: int32(170004),
			Stmts: []*pg_query.RawStmt{
				{
					Stmt: &pg_query.Node{
						Node: &pg_query.Node_CreateStmt{
							CreateStmt: &pg_query.CreateStmt{
								Relation: pg_query.MakeSimpleRangeVar("test", 13),
								TableElts: []*pg_query.Node{
									pg_query.MakeSimpleColumnDefNode(
										"id",
										&pg_query.TypeName{
											Names: []*pg_query.Node{
												pg_query.MakeStrNode("serial"),
											},
											Typemod:  -1,
											Location: 27,
										},
										[]*pg_query.Node{
											pg_query.MakePrimaryKeyConstraintNode(34),
										},
										24,
									),
									pg_query.MakeSimpleColumnDefNode(
										"user_id",
										&pg_query.TypeName{
											Names: []*pg_query.Node{
												pg_query.MakeStrNode("pg_catalog"),
												pg_query.MakeStrNode("int4"),
											},
											Typemod:  -1,
											Location: 59,
										},
										[]*pg_query.Node{
											pg_query.MakeDefaultConstraintNode(pg_query.MakeAConstIntNode(0, 75), 67),
											pg_query.MakeNotNullConstraintNode(77),
										},
										51,
									),
									pg_query.MakeSimpleColumnDefNode(
										"created_at",
										&pg_query.TypeName{
											Names: []*pg_query.Node{
												pg_query.MakeStrNode("pg_catalog"),
												pg_query.MakeStrNode("timestamp"),
											},
											Typemod:  -1,
											Location: 102,
										},
										[]*pg_query.Node{
											pg_query.MakeNotNullConstraintNode(130),
										},
										91,
									),
								},
								Oncommit: pg_query.OnCommitAction_ONCOMMIT_NOOP,
							},
						},
					},
					StmtLen: 139,
				},
			},
		},
	},
	{
		`SELECT * FROM a(1)`,
		`{"version":170004,"stmts":[{"stmt":{"SelectStmt":{"targetList":[{"ResTarget":{"val":{"ColumnRef":{"fields":[{"A_Star":{}}],"location":7}},"location":7}}],"fromClause":[{"RangeFunction":{"functions":[{"List":{"items":[{"FuncCall":{"funcname":[{"String":{"sval":"a"}}],"args":[{"A_Const":{"ival":{"ival":1},"location":16}}],"funcformat":"COERCE_EXPLICIT_CALL","location":14}},{}]}}]}}],"limitOption":"LIMIT_OPTION_DEFAULT","op":"SETOP_NONE"}}}]}`,
		&pg_query.ParseResult{
			Version: int32(170004),
			Stmts: []*pg_query.RawStmt{
				{
					Stmt: &pg_query.Node{
						Node: &pg_query.Node_SelectStmt{
							SelectStmt: &pg_query.SelectStmt{
								LimitOption: pg_query.LimitOption_LIMIT_OPTION_DEFAULT,
								Op:          pg_query.SetOperation_SETOP_NONE,
								TargetList: []*pg_query.Node{
									pg_query.MakeResTargetNodeWithVal(
										pg_query.MakeColumnRefNode(
											[]*pg_query.Node{
												pg_query.MakeAStarNode(),
											},
											7,
										),
										7,
									),
								},
								FromClause: []*pg_query.Node{
									pg_query.MakeSimpleRangeFunctionNode([]*pg_query.Node{
										pg_query.MakeListNode([]*pg_query.Node{
											pg_query.MakeFuncCallNode(
												[]*pg_query.Node{pg_query.MakeStrNode("a")},
												[]*pg_query.Node{pg_query.MakeAConstIntNode(1, 16)},
												14,
											),
											nil,
										}),
									}),
								},
							},
						},
					},
				},
			},
		},
	},
	{
		// Test for null-byte related crashes
		string([]byte{'S', 'E', 'L', 'E', 'C', 'T', ' ', '1', '\x00'}),
		`{"version":170004,"stmts":[{"stmt":{"SelectStmt":{"targetList":[{"ResTarget":{"val":{"A_Const":{"ival":{"ival":1},"location":7}},"location":7}}],"limitOption":"LIMIT_OPTION_DEFAULT","op":"SETOP_NONE"}}}]}`,
		&pg_query.ParseResult{
			Version: int32(170004),
			Stmts: []*pg_query.RawStmt{
				{
					Stmt: &pg_query.Node{
						Node: &pg_query.Node_SelectStmt{
							SelectStmt: &pg_query.SelectStmt{
								LimitOption: pg_query.LimitOption_LIMIT_OPTION_DEFAULT,
								Op:          pg_query.SetOperation_SETOP_NONE,
								TargetList: []*pg_query.Node{
									pg_query.MakeResTargetNodeWithVal(
										pg_query.MakeAConstIntNode(1, 7),
										7,
									),
								},
							},
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
			t.Errorf("protobuf error %s\n\n", err)
		} else if diff := cmp.Diff(actualTree, test.expectedTree, protocmp.Transform()); diff != "" {
			t.Errorf("protobuf unexpected difference:\n%v", diff)
		}
	}
}

var parseErrorTests = []struct {
	input       string
	expectedErr error
}{
	{
		"SELECT $",
		&parser.Error{
			Message:   "syntax error at or near \"$\"",
			Cursorpos: 8,
			Filename:  "scan.l",
			Funcname:  "scanner_yyerror",
		},
	},
	{
		"SELECT * FROM y WHERE x IN ($1, ",
		&parser.Error{
			Message:   "syntax error at end of input",
			Cursorpos: 33,
			Filename:  "scan.l",
			Funcname:  "scanner_yyerror",
		},
	},
}

func TestParseError(t *testing.T) {
	for _, test := range parseErrorTests {
		_, actualErr := pg_query.Parse(test.input)

		if actualErr == nil {
			t.Errorf("Parse(%s)\nexpected error but none returned\n\n", test.input)
		} else {
			exp := test.expectedErr.(*parser.Error)
			act := actualErr.(*parser.Error)
			act.Lineno = 0 // Line number is architecture dependent, so we ignore it
			if !reflect.DeepEqual(act, exp) {
				t.Errorf(
					"Parse(%s)\nexpected error %s at %d (%s:%d), func: %s, context: %s\nactual error %+v at %d (%s:%d), func: %s, context: %s\n\n",
					test.input,
					exp.Message, exp.Cursorpos, exp.Filename, exp.Lineno, exp.Funcname, exp.Context,
					act.Message, act.Cursorpos, act.Filename, act.Lineno, act.Funcname, act.Context)
			}
		}
	}
}

func TestParseConcurrency(t *testing.T) {
	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()

			_, err := pg_query.Parse("SELECT 1 FROM x WHERE y IN ('a', 'b', 'c')")

			if err != nil {
				t.Errorf("Concurrency test produced error %s\n\n", err)
			}

			fmt.Printf(".")
		}()
	}

	wg.Wait()
	fmt.Println()
}

var parsePlPgSQLTests = []struct {
	input        string
	expectedJSON string
}{
	{
		`CREATE OR REPLACE FUNCTION cs_fmt_browser_version(v_name varchar,` +
			`v_version varchar) ` +
			`RETURNS varchar AS $$ ` +
			`BEGIN ` +
			`    IF v_version IS NULL THEN` +
			`        RETURN v_name;` +
			`    END IF;` +
			`    RETURN v_name || '/' || v_version;` +
			`END;` +
			`$$ LANGUAGE plpgsql;`,
		`[
{"PLpgSQL_function":{"datums":[{"PLpgSQL_var":{"refname":"v_name","datatype":{"PLpgSQL_type":{"typname":"pg_catalog.\"varchar\""}}}},{"PLpgSQL_var":{"refname":"v_version","datatype":{"PLpgSQL_type":{"typname":"pg_catalog.\"varchar\""}}}},{"PLpgSQL_var":{"refname":"found","datatype":{"PLpgSQL_type":{"typname":"pg_catalog.\"boolean\""}}}}],"action":{"PLpgSQL_stmt_block":{"lineno":1,"body":[{"PLpgSQL_stmt_if":{"lineno":1,"cond":{"PLpgSQL_expr":{"query":"v_version IS NULL","parseMode":2}},"then_body":[{"PLpgSQL_stmt_return":{"lineno":1,"expr":{"PLpgSQL_expr":{"query":"v_name","parseMode":2}}}}]}},{"PLpgSQL_stmt_return":{"lineno":1,"expr":{"PLpgSQL_expr":{"query":"v_name || '/' || v_version","parseMode":2}}}}]}}}}
]`,
	},
}

func TestParsePlPgSQL(t *testing.T) {
	for _, test := range parsePlPgSQLTests {
		actualJSON, err := pg_query.ParsePlPgSqlToJSON(test.input)

		if err != nil {
			t.Errorf("ParsePlPgSqlToJSON(%s)\nerror %s\n\n", test.input, err)
		} else if actualJSON != test.expectedJSON {
			t.Errorf("ParsePlPgSqlToJSON(%s)\nexpected %s\nactual %s\n\n", test.input, test.expectedJSON, actualJSON)
		}
	}
}

func TestScan(t *testing.T) {
	smokeTest := func(input string) {
		_, err := pg_query.Scan(input)
		if err != nil {
			t.Error(err)
		}
	}
	for _, testCase := range parseTests {
		smokeTest(testCase.input)
	}
	for _, testCase := range parsePlPgSQLTests {
		smokeTest(testCase.input)
	}
}
