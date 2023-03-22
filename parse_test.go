package pg_query_test

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"sync"
	"testing"

	"github.com/google/go-cmp/cmp"
	pg_query "github.com/pganalyze/pg_query_go/v4"
	"github.com/pganalyze/pg_query_go/v4/parser"
	"google.golang.org/protobuf/testing/protocmp"
)

var parseTests = []struct {
	input        string
	expectedJSON string
	expectedTree *pg_query.ParseResult
}{
	{
		"SELECT 1",
		`{"version":150001,"stmts":[{"stmt":{"SelectStmt":{"targetList":[{"ResTarget":{"val":{"A_Const":{"ival":{"ival":1},"location":7}},"location":7}}],"limitOption":"LIMIT_OPTION_DEFAULT","op":"SETOP_NONE"}}}]}`,
		&pg_query.ParseResult{
			Version: int32(150001),
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
		`{"version":150001,"stmts":[{"stmt":{"SelectStmt":{"targetList":[{"ResTarget":{"val":{"ColumnRef":{"fields":[{"A_Star":{}}],"location":7}},"location":7}}],"fromClause":[{"RangeVar":{"relname":"x","inh":true,"relpersistence":"p","location":14}}],"whereClause":{"A_Expr":{"kind":"AEXPR_OP","name":[{"String":{"sval":"="}}],"lexpr":{"ColumnRef":{"fields":[{"String":{"sval":"z"}}],"location":22}},"rexpr":{"A_Const":{"ival":{"ival":1},"location":26}},"location":24}},"limitOption":"LIMIT_OPTION_DEFAULT","op":"SETOP_NONE"}}}]}`,
		&pg_query.ParseResult{
			Version: int32(150001),
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
		`{"version":150001,"stmts":[{"stmt":{"InsertStmt":{"relation":{"relname":"schema_index_stats","inh":true,"relpersistence":"p","location":12},"cols":[{"ResTarget":{"name":"snapshot_id","location":34}},{"ResTarget":{"name":"schema_index_id","location":48}},{"ResTarget":{"name":"size_bytes","location":66}}],"selectStmt":{"SelectStmt":{"valuesLists":[{"List":{"items":[{"A_Const":{"ival":{"ival":11710849},"location":88}},{"A_Const":{"ival":{"ival":8448632},"location":97}},{"A_Const":{"ival":{"ival":16384},"location":105}}]}},{"List":{"items":[{"A_Const":{"ival":{"ival":11710849},"location":113}},{"A_Const":{"ival":{"ival":8448633},"location":122}},{"A_Const":{"ival":{"ival":16384},"location":130}}]}}],"limitOption":"LIMIT_OPTION_DEFAULT","op":"SETOP_NONE"}},"returningList":[{"ResTarget":{"val":{"ColumnRef":{"fields":[{"String":{"sval":"id"}}],"location":147}},"location":147}}],"override":"OVERRIDING_NOT_SET"}}}]}`,
		&pg_query.ParseResult{
			Version: int32(150001),
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
		`{"version":150001,"stmts":[{"stmt":{"SelectStmt":{"targetList":[{"ResTarget":{"val":{"ColumnRef":{"fields":[{"A_Star":{}}],"location":7}},"location":7}}],"fromClause":[{"RangeVar":{"relname":"x","inh":true,"relpersistence":"p","location":14}}],"whereClause":{"A_Expr":{"kind":"AEXPR_IN","name":[{"String":{"sval":"="}}],"lexpr":{"ColumnRef":{"fields":[{"String":{"sval":"y"}}],"location":22}},"rexpr":{"List":{"items":[{"ParamRef":{"number":1,"location":28}}]}},"location":24}},"limitOption":"LIMIT_OPTION_DEFAULT","op":"SETOP_NONE"}}}]}`,
		&pg_query.ParseResult{
			Version: int32(150001),
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
		`{"version":150001,"stmts":[{"stmt":{"SelectStmt":{"targetList":[{"ResTarget":{"name":"Schema","val":{"ColumnRef":{"fields":[{"String":{"sval":"n"}},{"String":{"sval":"nspname"}}],"location":7}},"location":7}},{"ResTarget":{"name":"Name","val":{"ColumnRef":{"fields":[{"String":{"sval":"c"}},{"String":{"sval":"relname"}}],"location":36}},"location":36}},{"ResTarget":{"name":"Type","val":{"CaseExpr":{"arg":{"ColumnRef":{"fields":[{"String":{"sval":"c"}},{"String":{"sval":"relkind"}}],"location":68}},"args":[{"CaseWhen":{"expr":{"A_Const":{"sval":{"sval":"r"},"location":83}},"result":{"A_Const":{"sval":{"sval":"table"},"location":92}},"location":78}},{"CaseWhen":{"expr":{"A_Const":{"sval":{"sval":"v"},"location":105}},"result":{"A_Const":{"sval":{"sval":"view"},"location":114}},"location":100}},{"CaseWhen":{"expr":{"A_Const":{"sval":{"sval":"m"},"location":126}},"result":{"A_Const":{"sval":{"sval":"materialized view"},"location":135}},"location":121}},{"CaseWhen":{"expr":{"A_Const":{"sval":{"sval":"i"},"location":160}},"result":{"A_Const":{"sval":{"sval":"index"},"location":169}},"location":155}},{"CaseWhen":{"expr":{"A_Const":{"sval":{"sval":"S"},"location":182}},"result":{"A_Const":{"sval":{"sval":"sequence"},"location":191}},"location":177}},{"CaseWhen":{"expr":{"A_Const":{"sval":{"sval":"s"},"location":207}},"result":{"A_Const":{"sval":{"sval":"special"},"location":216}},"location":202}},{"CaseWhen":{"expr":{"A_Const":{"sval":{"sval":"f"},"location":231}},"result":{"A_Const":{"sval":{"sval":"foreign table"},"location":240}},"location":226}}],"location":63}},"location":63}},{"ResTarget":{"name":"Owner","val":{"FuncCall":{"funcname":[{"String":{"sval":"pg_catalog"}},{"String":{"sval":"pg_get_userbyid"}}],"args":[{"ColumnRef":{"fields":[{"String":{"sval":"c"}},{"String":{"sval":"relowner"}}],"location":304}}],"funcformat":"COERCE_EXPLICIT_CALL","location":277}},"location":277}}],"fromClause":[{"JoinExpr":{"jointype":"JOIN_LEFT","larg":{"RangeVar":{"schemaname":"pg_catalog","relname":"pg_class","inh":true,"relpersistence":"p","alias":{"aliasname":"c"},"location":336}},"rarg":{"RangeVar":{"schemaname":"pg_catalog","relname":"pg_namespace","inh":true,"relpersistence":"p","alias":{"aliasname":"n"},"location":374}},"quals":{"A_Expr":{"kind":"AEXPR_OP","name":[{"String":{"sval":"="}}],"lexpr":{"ColumnRef":{"fields":[{"String":{"sval":"n"}},{"String":{"sval":"oid"}}],"location":403}},"rexpr":{"ColumnRef":{"fields":[{"String":{"sval":"c"}},{"String":{"sval":"relnamespace"}}],"location":411}},"location":409}}}}],"whereClause":{"BoolExpr":{"boolop":"AND_EXPR","args":[{"A_Expr":{"kind":"AEXPR_IN","name":[{"String":{"sval":"="}}],"lexpr":{"ColumnRef":{"fields":[{"String":{"sval":"c"}},{"String":{"sval":"relkind"}}],"location":435}},"rexpr":{"List":{"items":[{"A_Const":{"sval":{"sval":"r"},"location":449}},{"A_Const":{"sval":{"sval":""},"location":453}}]}},"location":445}},{"A_Expr":{"kind":"AEXPR_OP","name":[{"String":{"sval":"\u003c\u003e"}}],"lexpr":{"ColumnRef":{"fields":[{"String":{"sval":"n"}},{"String":{"sval":"nspname"}}],"location":467}},"rexpr":{"A_Const":{"sval":{"sval":"pg_catalog"},"location":480}},"location":477}},{"A_Expr":{"kind":"AEXPR_OP","name":[{"String":{"sval":"\u003c\u003e"}}],"lexpr":{"ColumnRef":{"fields":[{"String":{"sval":"n"}},{"String":{"sval":"nspname"}}],"location":503}},"rexpr":{"A_Const":{"sval":{"sval":"information_schema"},"location":516}},"location":513}},{"A_Expr":{"kind":"AEXPR_OP","name":[{"String":{"sval":"!~"}}],"lexpr":{"ColumnRef":{"fields":[{"String":{"sval":"n"}},{"String":{"sval":"nspname"}}],"location":547}},"rexpr":{"A_Const":{"sval":{"sval":"^pg_toast"},"location":560}},"location":557}},{"FuncCall":{"funcname":[{"String":{"sval":"pg_catalog"}},{"String":{"sval":"pg_table_is_visible"}}],"args":[{"ColumnRef":{"fields":[{"String":{"sval":"c"}},{"String":{"sval":"oid"}}],"location":613}}],"funcformat":"COERCE_EXPLICIT_CALL","location":582}}],"location":463}},"sortClause":[{"SortBy":{"node":{"A_Const":{"ival":{"ival":1},"location":632}},"sortby_dir":"SORTBY_DEFAULT","sortby_nulls":"SORTBY_NULLS_DEFAULT","location":-1}},{"SortBy":{"node":{"A_Const":{"ival":{"ival":2},"location":634}},"sortby_dir":"SORTBY_DEFAULT","sortby_nulls":"SORTBY_NULLS_DEFAULT","location":-1}}],"limitOption":"LIMIT_OPTION_DEFAULT","op":"SETOP_NONE"}},"stmt_len":635}]}`,
		&pg_query.ParseResult{
			Version: int32(150001),
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
		`{"version":150001,"stmts":[{"stmt":{"CreateFunctionStmt":{"funcname":[{"String":{"sval":"change_trigger_v2"}}],"returnType":{"names":[{"String":{"sval":"trigger"}}],"typemod":-1,"location":44},"options":[{"DefElem":{"defname":"language","arg":{"String":{"sval":"plpgsql"}},"defaction":"DEFELEM_UNSPEC","location":53}},{"DefElem":{"defname":"as","arg":{"List":{"items":[{"String":{"sval":"\n\t\tDECLARE\n\t\tBEGIN\n\t\t\tPERFORM 'dummy';\n\t\tEND;\n\t\t"}}]}},"defaction":"DEFELEM_UNSPEC","location":71}}]}},"stmt_len":126}]}`,
		&pg_query.ParseResult{
			Version: int32(150001),
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
		`{"version":150001,"stmts":[{"stmt":{"CreateStmt":{"relation":{"relname":"test","inh":true,"relpersistence":"p","location":13},"tableElts":[{"ColumnDef":{"colname":"id","typeName":{"names":[{"String":{"sval":"serial"}}],"typemod":-1,"location":27},"is_local":true,"constraints":[{"Constraint":{"contype":"CONSTR_PRIMARY","location":34}}],"location":24}},{"ColumnDef":{"colname":"user_id","typeName":{"names":[{"String":{"sval":"pg_catalog"}},{"String":{"sval":"int4"}}],"typemod":-1,"location":59},"is_local":true,"constraints":[{"Constraint":{"contype":"CONSTR_DEFAULT","location":67,"raw_expr":{"A_Const":{"ival":{},"location":75}}}},{"Constraint":{"contype":"CONSTR_NOTNULL","location":77}}],"location":51}},{"ColumnDef":{"colname":"created_at","typeName":{"names":[{"String":{"sval":"pg_catalog"}},{"String":{"sval":"timestamp"}}],"typemod":-1,"location":102},"is_local":true,"constraints":[{"Constraint":{"contype":"CONSTR_NOTNULL","location":130}}],"location":91}}],"oncommit":"ONCOMMIT_NOOP"}},"stmt_len":139}]}`,
		&pg_query.ParseResult{
			Version: int32(150001),
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
		`{"version":150001,"stmts":[{"stmt":{"SelectStmt":{"targetList":[{"ResTarget":{"val":{"ColumnRef":{"fields":[{"A_Star":{}}],"location":7}},"location":7}}],"fromClause":[{"RangeFunction":{"functions":[{"List":{"items":[{"FuncCall":{"funcname":[{"String":{"sval":"a"}}],"args":[{"A_Const":{"ival":{"ival":1},"location":16}}],"funcformat":"COERCE_EXPLICIT_CALL","location":14}},{}]}}]}}],"limitOption":"LIMIT_OPTION_DEFAULT","op":"SETOP_NONE"}}}]}`,
		&pg_query.ParseResult{
			Version: int32(150001),
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
		`{"version":150001,"stmts":[{"stmt":{"SelectStmt":{"targetList":[{"ResTarget":{"val":{"A_Const":{"ival":{"ival":1},"location":7}},"location":7}}],"limitOption":"LIMIT_OPTION_DEFAULT","op":"SETOP_NONE"}}}]}`,
		&pg_query.ParseResult{
			Version: int32(150001),
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
{"PLpgSQL_function":{"datums":[{"PLpgSQL_var":{"refname":"v_name","datatype":{"PLpgSQL_type":{"typname":"UNKNOWN"}}}},{"PLpgSQL_var":{"refname":"v_version","datatype":{"PLpgSQL_type":{"typname":"UNKNOWN"}}}},{"PLpgSQL_var":{"refname":"found","datatype":{"PLpgSQL_type":{"typname":"UNKNOWN"}}}}],"action":{"PLpgSQL_stmt_block":{"lineno":1,"body":[{"PLpgSQL_stmt_if":{"lineno":1,"cond":{"PLpgSQL_expr":{"query":"v_version IS NULL"}},"then_body":[{"PLpgSQL_stmt_return":{"lineno":1,"expr":{"PLpgSQL_expr":{"query":"v_name"}}}}]}},{"PLpgSQL_stmt_return":{"lineno":1,"expr":{"PLpgSQL_expr":{"query":"v_name || '/' || v_version"}}}}]}}}}
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
func strPtr(s string) *string {
	return &s
}
func S2PtrString(a [2]*string) (s string) {
	if a[0] == nil {
		s += fmt.Sprintf("{nil, ")
	} else {
		s += fmt.Sprintf("{%s, ", *a[0])
	}
	if a[1] == nil {
		s += fmt.Sprintf("nil}")
	} else {
		s += fmt.Sprintf("%s}", *a[1])
	}
	return s
}
func TestFilterColumns(t *testing.T) {
	inputs := []string{
		"SELECT * FROM x WHERE y = $1 AND z = 1",
		"SELECT * FROM x WHERE x.y = $1 AND x.z = 1",
		"WITH a AS (SELECT * FROM x WHERE x.y = $1 AND x.z = 1) SELECT * FROM a WHERE b = 5",
		"SELECT * FROM x WHERE x.y IS TRUE AND x.z IS NOT FALSE",
		"SELECT * FROM x WHERE x.y = COALESCE(z.a, z.b)",
		"CREATE INDEX testidx ON test USING btree (a, (lower(b) || upper(c))) WHERE pow(a, 2) > 25",
		"SELECT * FROM pg_catalog.pg_class c JOIN (SELECT 17650 AS oid UNION ALL SELECT 17663 AS oid) vals ON c.oid = vals.oid",
	}
	combiners := []string{
		"UNION", "UNION ALL", "EXCEPT", "EXCEPT ALL", "INTERSECT", "INTERSECT ALL",
	}
	expected := [][][2]*string{[][2]*string{[2]*string{nil, strPtr("y")}, [2]*string{nil, strPtr("z")}}, // finds unqualified names
		[][2]*string{[2]*string{strPtr("x"), strPtr("y")}, [2]*string{strPtr("x"), strPtr("z")}},                                       // finds qualified names
		[][2]*string{[2]*string{nil, strPtr("b")}, [2]*string{strPtr("x"), strPtr("y")}, [2]*string{strPtr("x"), strPtr("z")}},         // traverses into CTEs
		[][2]*string{[2]*string{strPtr("x"), strPtr("y")}, [2]*string{strPtr("x"), strPtr("z")}},                                       // recognizes boolean tests
		[][2]*string{[2]*string{strPtr("x"), strPtr("y")}, [2]*string{strPtr("z"), strPtr("a")}, [2]*string{strPtr("z"), strPtr("b")}}, // finds unqualified names in #{combiner} query
		[][2]*string{[2]*string{nil, strPtr("a")}},
		[][2]*string{[2]*string{strPtr("pg_catalog.pg_class"), strPtr("oid")}, [2]*string{strPtr("vals"), strPtr("oid")}},
	}

	for _, input := range combiners {
		inputs = append(inputs, fmt.Sprintf("SELECT * FROM x where y = $1 %s SELECT * FROM x where z = $2", input))
		expected = append(expected, [][2]*string{[2]*string{nil, strPtr("y")}, [2]*string{nil, strPtr("z")}})
	}
	for i, e := range inputs {
		pr, err := pg_query.Parse(e)
		if err != nil {
			t.Errorf("Parse() failed: %s", err)
			continue
		}
		tpr := pg_query.NewTaggedParseResult(pr)
		if err := tpr.LoadObjects(); err != nil {
			t.Errorf("LoadObjects() failed: %s", err)
			continue
		}
		filtersColumns, err := tpr.FilterColumns()
		if err != nil {
			t.Errorf("FilterColumns() failed: %s", err)
			continue
		}
		if len(filtersColumns) != len(expected[i]) {
			t.Errorf("FilterColumns() failed: input: %s, expected %d, actual %d", e, len(expected), len(filtersColumns))
			continue
		}
		for j, f := range filtersColumns {
			s1 := expected[i][j][0]
			s2 := expected[i][j][1]
			if (s1 == nil && f.Schema != nil) || (s2 == nil && f.Rel != nil) {
				t.Errorf("FilterColumns() failed: input: %s, expected %v, actual %v", e, S2PtrString(expected[i][j]), f)
				continue
			}
			if s1 != nil {
				if f.Schema == nil {
					t.Errorf("FilterColumns() failed: input: %s, expected %v, actual %v", e, S2PtrString(expected[i][j]), f)

					continue
				}
				if *(f.Schema) != *s1 {
					t.Errorf("FilterColumns() failed: input: %s, expected %v, actual %v", e, S2PtrString(expected[i][j]), f)
					continue
				}
			}
			if s2 != nil {
				if f.Rel == nil {
					t.Errorf("FilterColumns() failed: input: %s, expected %v, actual %v", e, S2PtrString(expected[i][j]), f)
					continue
				}
				if *(f.Rel) != *s2 {
					t.Errorf("FilterColumns() failed: input: %s, expected %v, actual %v", e, S2PtrString(expected[i][j]), f)
					continue
				}
			}
		}
	}
}

func mulStringHlp(pre, post []string, format string, n int) *[]string {
	var ret []string
	for i := 0; i < (n - 1); i++ {
		ret = append(ret, fmt.Sprintf(format, n-i))
	}
	ret = append(pre, ret...)
	ret = append(ret, post...)
	return &ret
}

func mulStringJoinFoo(n int) string {
	var buf []string
	for i := 1; i <= n; i++ {
		buf = append(buf, fmt.Sprintf("JOIN foo_%d ON foo.id = foo_%d.foo_id", i, i))
	}
	return strings.Join(buf, " ")
}

func TestLoadObjects(t *testing.T) {

	type expectedOutput struct {
		Query     string
		Tables    *[]string
		Functions *[]string
		FuncArgs  *[]string
		Aliases   *map[string]string
		CteNames  *[]string
		Err       error
	}
	/*
		SELECT DISTINCT\n     base_src.oid AS base_tbl_oid\n     ,srcobj.oid as child_id\n     ,srcnsp.nspname AS child_schema\n     ,srcobj.relname AS child_name\n     ,case when srcobj.oid=base_src.oid then 'tmp_site_questionnaire_response__c_550c2a6'\n     else 'tmp_' + srcobj.relname + '_57ded757d06e4933875e65b7a4423752' end as child_tmp_name\n     ,tgtobj.oid AS parent_id\n     ,tgtnsp.nspname AS parent_schema\n     ,tgtobj.relname AS parent_name\n     ,'tmp_' + tgtobj.relname + '_57ded757d06e4933875e65b7a4423752' as parent_tmp_name\n     ,'old_' + tgtobj.relname + '_57ded757d06e4933875e65b7a4423752' as parent_old_name\n FROM\n     pg_catalog.pg_class AS srcobj\n INNER JOIN\n     pg_catalog.pg_depend AS srcdep\n         ON srcobj.oid = srcdep.refobjid\n INNER JOIN\n     pg_catalog.pg_depend AS tgtdep\n         ON srcdep.objid = tgtdep.objid\n JOIN\n     pg_catalog.pg_class AS tgtobj\n         ON tgtdep.refobjid = tgtobj.oid AND srcobj.oid <> tgtobj.oid\n LEFT OUTER JOIN\n     pg_catalog.pg_namespace AS srcnsp\n         ON srcobj.relnamespace = srcnsp.oid\n LEFT OUTER JOIN\n     pg_catalog.pg_namespace tgtnsp\n         ON tgtobj.relnamespace = tgtnsp.oid\n CROSS JOIN (SELECT srcobj.oid\n             FROM  pg_catalog.pg_class AS srcobj\n             INNER JOIN pg_catalog.pg_namespace AS srcnsp\n             ON srcobj.relnamespace = srcnsp.oid\n             WHERE relname='site_questionnaire_response__c'\n             and srcnsp.nspname = 'salesforce') base_src\n WHERE tgtdep.deptype = 'i' \n AND tgtobj.relkind = 'v' \n AND srcnsp.nspname = 'salesforce'
	*/
	expectations := []expectedOutput{

		{
			Query:     "SELECT * FROM foo.testfunc(test_Table);",
			Tables:    &[]string{},
			Functions: &[]string{"foo.testfunc"},
			FuncArgs:  &[]string{"test_table"},
			Aliases:   &map[string]string{},
			CteNames:  &[]string{}},
		{
			Query:     "SELECT a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(a(b))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))",
			Tables:    &[]string{},
			Functions: &[]string{"a"},
			FuncArgs:  &[]string{},
			Aliases:   &map[string]string{},
			CteNames:  &[]string{}},
		{
			Query:     "SELECT * FROM \"t0\"\n    JOIN \"t1\" ON (1) JOIN \"t2\" ON (1) JOIN \"t3\" ON (1) JOIN \"t4\" ON (1) JOIN \"t5\" ON (1)\n    JOIN \"t6\" ON (1) JOIN \"t7\" ON (1) JOIN \"t8\" ON (1) JOIN \"t9\" ON (1) JOIN \"t10\" ON (1)\n    JOIN \"t11\" ON (1) JOIN \"t12\" ON (1) JOIN \"t13\" ON (1) JOIN \"t14\" ON (1) JOIN \"t15\" ON (1)\n    JOIN \"t16\" ON (1) JOIN \"t17\" ON (1) JOIN \"t18\" ON (1) JOIN \"t19\" ON (1) JOIN \"t20\" ON (1)\n    JOIN \"t21\" ON (1) JOIN \"t22\" ON (1) JOIN \"t23\" ON (1) JOIN \"t24\" ON (1) JOIN \"t25\" ON (1)\n    JOIN \"t26\" ON (1) JOIN \"t27\" ON (1) JOIN \"t28\" ON (1) JOIN \"t29\" ON (1)",
			Tables:    &[]string{"t29", "t28", "t27", "t26", "t25", "t24", "t23", "t22", "t21", "t20", "t19", "t18", "t17", "t16", "t15", "t14", "t13", "t12", "t11", "t10", "t9", "t8", "t7", "t6", "t5", "t4", "t3", "t2", "t0", "t1"},
			Functions: &[]string{},
			FuncArgs:  &[]string{},
			Aliases:   &map[string]string{},
			CteNames:  &[]string{}},
		{
			Query:     "SELECT * FROM foo JOIN foo_1 ON foo.id = foo_1.foo_id JOIN foo_2 ON foo.id = foo_2.foo_id JOIN foo_3 ON foo.id = foo_3.foo_id JOIN foo_4 ON foo.id = foo_4.foo_id JOIN foo_5 ON foo.id = foo_5.foo_id JOIN foo_6 ON foo.id = foo_6.foo_id JOIN foo_7 ON foo.id = foo_7.foo_id JOIN foo_8 ON foo.id = foo_8.foo_id JOIN foo_9 ON foo.id = foo_9.foo_id JOIN foo_10 ON foo.id = foo_10.foo_id JOIN foo_11 ON foo.id = foo_11.foo_id JOIN foo_12 ON foo.id = foo_12.foo_id JOIN foo_13 ON foo.id = foo_13.foo_id JOIN foo_14 ON foo.id = foo_14.foo_id JOIN foo_15 ON foo.id = foo_15.foo_id JOIN foo_16 ON foo.id = foo_16.foo_id JOIN foo_17 ON foo.id = foo_17.foo_id JOIN foo_18 ON foo.id = foo_18.foo_id JOIN foo_19 ON foo.id = foo_19.foo_id JOIN foo_20 ON foo.id = foo_20.foo_id JOIN foo_21 ON foo.id = foo_21.foo_id JOIN foo_22 ON foo.id = foo_22.foo_id JOIN foo_23 ON foo.id = foo_23.foo_id JOIN foo_24 ON foo.id = foo_24.foo_id JOIN foo_25 ON foo.id = foo_25.foo_id JOIN foo_26 ON foo.id = foo_26.foo_id JOIN foo_27 ON foo.id = foo_27.foo_id JOIN foo_28 ON foo.id = foo_28.foo_id JOIN foo_29 ON foo.id = foo_29.foo_id JOIN foo_30 ON foo.id = foo_30.foo_id JOIN foo_31 ON foo.id = foo_31.foo_id JOIN foo_32 ON foo.id = foo_32.foo_id JOIN foo_33 ON foo.id = foo_33.foo_id JOIN foo_34 ON foo.id = foo_34.foo_id JOIN foo_35 ON foo.id = foo_35.foo_id JOIN foo_36 ON foo.id = foo_36.foo_id JOIN foo_37 ON foo.id = foo_37.foo_id JOIN foo_38 ON foo.id = foo_38.foo_id JOIN foo_39 ON foo.id = foo_39.foo_id JOIN foo_40 ON foo.id = foo_40.foo_id JOIN foo_41 ON foo.id = foo_41.foo_id JOIN foo_42 ON foo.id = foo_42.foo_id JOIN foo_43 ON foo.id = foo_43.foo_id JOIN foo_44 ON foo.id = foo_44.foo_id JOIN foo_45 ON foo.id = foo_45.foo_id JOIN foo_46 ON foo.id = foo_46.foo_id JOIN foo_47 ON foo.id = foo_47.foo_id JOIN foo_48 ON foo.id = foo_48.foo_id JOIN foo_49 ON foo.id = foo_49.foo_id JOIN foo_50 ON foo.id = foo_50.foo_id JOIN foo_51 ON foo.id = foo_51.foo_id JOIN foo_52 ON foo.id = foo_52.foo_id JOIN foo_53 ON foo.id = foo_53.foo_id JOIN foo_54 ON foo.id = foo_54.foo_id JOIN foo_55 ON foo.id = foo_55.foo_id JOIN foo_56 ON foo.id = foo_56.foo_id JOIN foo_57 ON foo.id = foo_57.foo_id JOIN foo_58 ON foo.id = foo_58.foo_id JOIN foo_59 ON foo.id = foo_59.foo_id JOIN foo_60 ON foo.id = foo_60.foo_id JOIN foo_61 ON foo.id = foo_61.foo_id JOIN foo_62 ON foo.id = foo_62.foo_id JOIN foo_63 ON foo.id = foo_63.foo_id JOIN foo_64 ON foo.id = foo_64.foo_id JOIN foo_65 ON foo.id = foo_65.foo_id JOIN foo_66 ON foo.id = foo_66.foo_id JOIN foo_67 ON foo.id = foo_67.foo_id JOIN foo_68 ON foo.id = foo_68.foo_id JOIN foo_69 ON foo.id = foo_69.foo_id JOIN foo_70 ON foo.id = foo_70.foo_id JOIN foo_71 ON foo.id = foo_71.foo_id JOIN foo_72 ON foo.id = foo_72.foo_id JOIN foo_73 ON foo.id = foo_73.foo_id JOIN foo_74 ON foo.id = foo_74.foo_id JOIN foo_75 ON foo.id = foo_75.foo_id JOIN foo_76 ON foo.id = foo_76.foo_id JOIN foo_77 ON foo.id = foo_77.foo_id JOIN foo_78 ON foo.id = foo_78.foo_id JOIN foo_79 ON foo.id = foo_79.foo_id JOIN foo_80 ON foo.id = foo_80.foo_id JOIN foo_81 ON foo.id = foo_81.foo_id JOIN foo_82 ON foo.id = foo_82.foo_id JOIN foo_83 ON foo.id = foo_83.foo_id JOIN foo_84 ON foo.id = foo_84.foo_id JOIN foo_85 ON foo.id = foo_85.foo_id JOIN foo_86 ON foo.id = foo_86.foo_id JOIN foo_87 ON foo.id = foo_87.foo_id JOIN foo_88 ON foo.id = foo_88.foo_id JOIN foo_89 ON foo.id = foo_89.foo_id JOIN foo_90 ON foo.id = foo_90.foo_id JOIN foo_91 ON foo.id = foo_91.foo_id JOIN foo_92 ON foo.id = foo_92.foo_id JOIN foo_93 ON foo.id = foo_93.foo_id JOIN foo_94 ON foo.id = foo_94.foo_id JOIN foo_95 ON foo.id = foo_95.foo_id JOIN foo_96 ON foo.id = foo_96.foo_id JOIN foo_97 ON foo.id = foo_97.foo_id JOIN foo_98 ON foo.id = foo_98.foo_id JOIN foo_99 ON foo.id = foo_99.foo_id JOIN foo_100 ON foo.id = foo_100.foo_id",
			Tables:    &[]string{"foo_100", "foo_99", "foo_98", "foo_97", "foo_96", "foo_95", "foo_94", "foo_93", "foo_92", "foo_91", "foo_90", "foo_89", "foo_88", "foo_87", "foo_86", "foo_85", "foo_84", "foo_83", "foo_82", "foo_81", "foo_80", "foo_79", "foo_78", "foo_77", "foo_76", "foo_75", "foo_74", "foo_73", "foo_72", "foo_71", "foo_70", "foo_69", "foo_68", "foo_67", "foo_66", "foo_65", "foo_64", "foo_63", "foo_62", "foo_61", "foo_60", "foo_59", "foo_58", "foo_57", "foo_56", "foo_55", "foo_54", "foo_53", "foo_52", "foo_51", "foo_50", "foo_49", "foo_48", "foo_47", "foo_46", "foo_45", "foo_44", "foo_43", "foo_42", "foo_41", "foo_40", "foo_39", "foo_38", "foo_37", "foo_36", "foo_35", "foo_34", "foo_33", "foo_32", "foo_31", "foo_30", "foo_29", "foo_28", "foo_27", "foo_26", "foo_25", "foo_24", "foo_23", "foo_22", "foo_21", "foo_20", "foo_19", "foo_18", "foo_17", "foo_16", "foo_15", "foo_14", "foo_13", "foo_12", "foo_11", "foo_10", "foo_9", "foo_8", "foo_7", "foo_6", "foo_5", "foo_4", "foo_3", "foo_2", "foo", "foo_1"},
			Functions: &[]string{},
			FuncArgs:  &[]string{},
			Aliases:   &map[string]string{},
			CteNames:  &[]string{}},
		{
			Query:     "SELECT memory_total_bytes, memory_free_bytes, memory_pagecache_bytes, memory_buffers_bytes, memory_applications_bytes, (memory_swap_total_bytes - memory_swap_free_bytes) AS swap, date_part($0, s.collected_at) AS collected_at FROM snapshots s JOIN system_snapshots ON (snapshot_id = s.id) WHERE s.database_id = $0 AND s.collected_at BETWEEN $0 AND $0 ORDER BY collected_at",
			Tables:    &[]string{"snapshots", "system_snapshots"},
			Functions: &[]string{"date_part"},
			FuncArgs:  &[]string{"snapshots.collected_at"},
			Aliases:   &map[string]string{"s": "snapshots"},
			CteNames:  &[]string{}},
		{
			Query:     "-- nothing",
			Tables:    &[]string{},
			Functions: &[]string{},
			FuncArgs:  &[]string{},
			Aliases:   &map[string]string{},
			CteNames:  &[]string{}},
		{
			Query:     "ALTER TABLE test ADD PRIMARY KEY (gid)",
			Tables:    &[]string{"test"},
			Functions: &[]string{},
			FuncArgs:  &[]string{},
			Aliases:   &map[string]string{},
			CteNames:  &[]string{}},
		{
			Query:     "SET statement_timeout=0",
			Tables:    &[]string{},
			Functions: &[]string{},
			FuncArgs:  &[]string{},
			Aliases:   &map[string]string{},
			CteNames:  &[]string{}},
		{
			Query:     "SHOW work_mem",
			Tables:    &[]string{},
			Functions: &[]string{},
			FuncArgs:  &[]string{},
			Aliases:   &map[string]string{},
			CteNames:  &[]string{}},
		{
			Query:     "COPY test (id) TO stdout",
			Tables:    &[]string{"test"},
			Functions: &[]string{},
			FuncArgs:  &[]string{},
			Aliases:   &map[string]string{},
			CteNames:  &[]string{}},
		{
			Query:     "drop table abc.test123 cascade",
			Tables:    &[]string{"abc.test123"},
			Functions: &[]string{},
			FuncArgs:  &[]string{},
			Aliases:   &map[string]string{},
			CteNames:  &[]string{}},
		{
			Query:     "COMMIT",
			Tables:    &[]string{},
			Functions: &[]string{},
			FuncArgs:  &[]string{},
			Aliases:   &map[string]string{},
			CteNames:  &[]string{}},
		{
			Query:     "CHECKPOINT",
			Tables:    &[]string{},
			Functions: &[]string{},
			FuncArgs:  &[]string{},
			Aliases:   &map[string]string{},
			CteNames:  &[]string{}},
		{
			Query:     "VACUUM my_table",
			Tables:    &[]string{"my_table"},
			Functions: &[]string{},
			FuncArgs:  &[]string{},
			Aliases:   &map[string]string{},
			CteNames:  &[]string{}},
		{
			Query:     "EXPLAIN DELETE FROM test",
			Tables:    &[]string{"test"},
			Functions: &[]string{},
			FuncArgs:  &[]string{},
			Aliases:   &map[string]string{},
			CteNames:  &[]string{}},
		{
			Query:     "CREATE TEMP TABLE test AS SELECT 1",
			Tables:    &[]string{"test"},
			Functions: &[]string{},
			FuncArgs:  &[]string{},
			Aliases:   &map[string]string{},
			CteNames:  &[]string{}},
		{
			Query:     "LOCK TABLE public.schema_migrations IN ACCESS SHARE MODE",
			Tables:    &[]string{"public.schema_migrations"},
			Functions: &[]string{},
			FuncArgs:  &[]string{},
			Aliases:   &map[string]string{},
			CteNames:  &[]string{}},
		{
			Query:     "CREATE TABLE test (a int4)",
			Tables:    &[]string{"test"},
			Functions: &[]string{},
			FuncArgs:  &[]string{},
			Aliases:   &map[string]string{},
			CteNames:  &[]string{}},
		{
			Query: "CREATE TABLE test (a int4) WITH OIDS",
			Err:   errors.New("syntax error at or near \"OIDS\"")},
		{
			Query:     "CREATE INDEX testidx ON test USING btree (a, (lower(b) || upper(c))) WHERE pow(a, 2) > 25",
			Tables:    &[]string{"test"},
			Functions: &[]string{"lower", "upper", "pow"},
			FuncArgs:  &[]string{"a"},
			Aliases:   &map[string]string{},
			CteNames:  &[]string{}},
		{
			Query:     "CREATE SCHEMA IF NOT EXISTS test AUTHORIZATION joe",
			Tables:    &[]string{},
			Functions: &[]string{},
			FuncArgs:  &[]string{},
			Aliases:   &map[string]string{},
			CteNames:  &[]string{}},
		{
			Query:     "CREATE VIEW myview AS SELECT * FROM mytab",
			Tables:    &[]string{"myview", "mytab"},
			Functions: &[]string{},
			FuncArgs:  &[]string{},
			Aliases:   &map[string]string{},
			CteNames:  &[]string{}},
		{
			Query:     "REFRESH MATERIALIZED VIEW myview",
			Tables:    &[]string{"myview"},
			Functions: &[]string{},
			FuncArgs:  &[]string{},
			Aliases:   &map[string]string{},
			CteNames:  &[]string{}},
		{
			Query:     "CREATE RULE shoe_ins_protect AS ON INSERT TO shoe\n                           DO INSTEAD NOTHING",
			Tables:    &[]string{"shoe"},
			Functions: &[]string{},
			FuncArgs:  &[]string{},
			Aliases:   &map[string]string{},
			CteNames:  &[]string{}},
		{
			Query:     "CREATE TRIGGER check_update\n                           BEFORE UPDATE ON accounts\n                           FOR EACH ROW\n                           EXECUTE PROCEDURE check_account_update()",
			Tables:    &[]string{"accounts"},
			Functions: &[]string{},
			FuncArgs:  &[]string{},
			Aliases:   &map[string]string{},
			CteNames:  &[]string{}},
		{
			Query:     "DROP SCHEMA myschema",
			Tables:    &[]string{},
			Functions: &[]string{},
			FuncArgs:  &[]string{},
			Aliases:   &map[string]string{},
			CteNames:  &[]string{}},
		{
			Query:     "DROP VIEW myview, myview2",
			Tables:    &[]string{},
			Functions: &[]string{},
			FuncArgs:  &[]string{},
			Aliases:   &map[string]string{},
			CteNames:  &[]string{}},
		{
			Query:     "DROP INDEX CONCURRENTLY myindex",
			Tables:    &[]string{},
			Functions: &[]string{},
			FuncArgs:  &[]string{},
			Aliases:   &map[string]string{},
			CteNames:  &[]string{}},
		{
			Query:     "DROP RULE myrule ON mytable CASCADE",
			Tables:    &[]string{"mytable"},
			Functions: &[]string{},
			FuncArgs:  &[]string{},
			Aliases:   &map[string]string{},
			CteNames:  &[]string{}},
		{
			Query:     "DROP TRIGGER IF EXISTS mytrigger ON mytable RESTRICT",
			Tables:    &[]string{"mytable"},
			Functions: &[]string{},
			FuncArgs:  &[]string{},
			Aliases:   &map[string]string{},
			CteNames:  &[]string{}},
		{
			Query:     "GRANT INSERT, UPDATE ON mytable TO myuser",
			Tables:    &[]string{"mytable"},
			Functions: &[]string{},
			FuncArgs:  &[]string{},
			Aliases:   &map[string]string{},
			CteNames:  &[]string{}},
		{
			Query:     "REVOKE admins FROM joe",
			Tables:    &[]string{},
			Functions: &[]string{},
			FuncArgs:  &[]string{},
			Aliases:   &map[string]string{},
			CteNames:  &[]string{}},
		{
			Query:     "TRUNCATE bigtable, \"fattable\" RESTART IDENTITY",
			Tables:    &[]string{"bigtable", "fattable"},
			Functions: &[]string{},
			FuncArgs:  &[]string{},
			Aliases:   &map[string]string{},
			CteNames:  &[]string{}},
		{
			Query:     "WITH a AS (SELECT * FROM x WHERE x.y = $1 AND x.z = 1) SELECT * FROM a",
			Tables:    &[]string{"x"},
			Functions: &[]string{},
			FuncArgs:  &[]string{},
			Aliases:   &map[string]string{},
			CteNames:  &[]string{"a"}},
		{
			Query:     "CREATE OR REPLACE FUNCTION thing(parameter_thing text)\n  RETURNS bigint AS\n$BODY$\nDECLARE\n        local_thing_id BIGINT := 0;\nBEGIN\n        SELECT thing_id INTO local_thing_id FROM thing_map\n        WHERE\n                thing_map_field = parameter_thing\n        ORDER BY 1 LIMIT 1;\n\n        IF NOT FOUND THEN\n                local_thing_id = 0;\n        END IF;\n        RETURN local_thing_id;\nEND;\n$BODY$\n  LANGUAGE plpgsql STABLE",
			Tables:    &[]string{},
			Functions: &[]string{"thing"},
			FuncArgs:  &[]string{},
			Aliases:   &map[string]string{},
			CteNames:  &[]string{}},
		{
			Query:     "CREATE FUNCTION getfoo(int) RETURNS TABLE (f1 int) AS '\n    SELECT * FROM foo WHERE fooid = $1;\n' LANGUAGE SQL",
			Tables:    &[]string{},
			Functions: &[]string{"getfoo"},
			FuncArgs:  &[]string{},
			Aliases:   &map[string]string{},
			CteNames:  &[]string{}},
		{
			Query:     "      CREATE OR REPLACE FUNCTION foo.testfunc(x integer) RETURNS integer AS $$\n      BEGIN\n      RETURN x\n      END;\n      $$ LANGUAGE plpgsql STABLE;\n",
			Tables:    &[]string{},
			Functions: &[]string{"foo.testfunc"},
			FuncArgs:  &[]string{},
			Aliases:   &map[string]string{},
			CteNames:  &[]string{}},
		{
			Query:     "      DROP FUNCTION IF EXISTS foo.testfunc(x integer);\n",
			Tables:    &[]string{},
			Functions: &[]string{"foo.testfunc"},
			FuncArgs:  &[]string{},
			Aliases:   &map[string]string{},
			CteNames:  &[]string{}},
		{
			Query:     "      ALTER FUNCTION foo.testfunc(integer) RENAME TO testfunc2;\n",
			Tables:    &[]string{},
			Functions: &[]string{"foo.testfunc", "testfunc2"},
			FuncArgs:  &[]string{},
			Aliases:   &map[string]string{},
			CteNames:  &[]string{}},
		{
			Query:     "select u.email, (select count(*) from enrollments e where e.user_id = u.id) as num_enrollments from users u",
			Tables:    &[]string{"users", "enrollments"},
			Functions: &[]string{"count"},
			FuncArgs:  &[]string{},
			Aliases:   &map[string]string{"u": "users", "e": "enrollments"},
			CteNames:  &[]string{}},
		{
			Query:     "WITH cte_name AS (SELECT 1) SELECT * FROM table_name, cte_name",
			Tables:    &[]string{"table_name"},
			Functions: &[]string{},
			FuncArgs:  &[]string{},
			Aliases:   &map[string]string{},
			CteNames:  &[]string{"cte_name"}},
		{
			Query:     "select u.* from (select * from users) u",
			Tables:    &[]string{"users"},
			Functions: &[]string{},
			FuncArgs:  &[]string{},
			Aliases:   &map[string]string{},
			CteNames:  &[]string{}},
		{
			Query:     "select users.id from users where 1 = (select count(*) from user_roles)",
			Tables:    &[]string{"users", "user_roles"},
			Functions: &[]string{"count"},
			FuncArgs:  &[]string{},
			Aliases:   &map[string]string{},
			CteNames:  &[]string{}},
		{
			Query:     "SELECT * FROM pg_catalog.pg_class c JOIN (SELECT 17650 AS oid UNION ALL SELECT 17663 AS oid) vals ON c.oid = vals.oid",
			Tables:    &[]string{"pg_catalog.pg_class"},
			Functions: &[]string{},
			FuncArgs:  &[]string{},
			Aliases:   &map[string]string{"c": "pg_catalog.pg_class"},
			CteNames:  &[]string{}},
		{
			Query:     "    select users.*\n    from users\n    where users.id IN (\n      select user_roles.user_id\n      from user_roles\n    ) and (users.created_at between '2016-06-01' and '2016-06-30')\n",
			Tables:    &[]string{"users", "user_roles"},
			Functions: &[]string{},
			FuncArgs:  &[]string{},
			Aliases:   &map[string]string{},
			CteNames:  &[]string{}},
		{
			Query:     "    select users.*\n    from users\n    order by (\n      select max(user_roles.role_id)\n      from user_roles\n      where user_roles.user_id = users.id\n    )\n",
			Tables:    &[]string{"users", "user_roles"},
			Functions: &[]string{"max"},
			FuncArgs:  &[]string{"user_roles.role_id"},
			Aliases:   &map[string]string{},
			CteNames:  &[]string{}},
		{
			Query:     "    select users.*\n    from users\n    order by (\n      select max(user_roles.role_id)\n      from user_roles\n      where user_roles.user_id = users.id\n    ) asc, (\n      select max(user_logins.role_id)\n      from user_logins\n      where user_logins.user_id = users.id\n    ) desc\n",
			Tables:    &[]string{"users", "user_roles", "user_logins"},
			Functions: &[]string{"max"},
			FuncArgs:  &[]string{"user_logins.role_id"},
			Aliases:   &map[string]string{},
			CteNames:  &[]string{}},
		{
			Query:     "    select users.*\n    from users\n    group by (\n      select max(user_roles.role_id)\n      from user_roles\n      where user_roles.user_id = users.id\n    )\n",
			Tables:    &[]string{"users", "user_roles"},
			Functions: &[]string{"max"},
			FuncArgs:  &[]string{"user_roles.role_id"},
			Aliases:   &map[string]string{},
			CteNames:  &[]string{}},
		{
			Query:     "    select users.*\n    from users\n    group by (\n      select max(user_roles.role_id)\n      from user_roles\n      where user_roles.user_id = users.id\n    ), (\n      select max(user_logins.role_id)\n      from user_logins\n      where user_logins.user_id = users.id\n    )\n",
			Tables:    &[]string{"users", "user_roles", "user_logins"},
			Functions: &[]string{"max"},
			FuncArgs:  &[]string{"user_logins.role_id"}, // FIXME SHOULD HAVE user_logins as well
			Aliases:   &map[string]string{},
			CteNames:  &[]string{}},
		{
			Query:     "    select users.*\n    from users\n    group by users.id\n    having 1 > (\n      select count(user_roles.role_id)\n      from user_roles\n      where user_roles.user_id = users.id\n    )\n",
			Tables:    &[]string{"users", "user_roles"},
			Functions: &[]string{"count"},
			FuncArgs:  &[]string{"user_roles.role_id"},
			Aliases:   &map[string]string{},
			CteNames:  &[]string{}},
		{
			Query:     "    select users.*\n    from users\n    group by users.id\n    having true and 1 > (\n      select count(user_roles.role_id)\n      from user_roles\n      where user_roles.user_id = users.id\n    )\n",
			Tables:    &[]string{"users", "user_roles"},
			Functions: &[]string{"count"},
			FuncArgs:  &[]string{"user_roles.role_id"},
			Aliases:   &map[string]string{},
			CteNames:  &[]string{}},
		{
			Query:     "    select foo.*\n    from foo\n    join ( select * from bar ) b\n    on b.baz = foo.quux\n",
			Tables:    &[]string{"foo", "bar"},
			Functions: &[]string{},
			FuncArgs:  &[]string{},
			Aliases:   &map[string]string{},
			CteNames:  &[]string{}},
		{
			Query:     "    SELECT *\n    FROM foo\n    INNER JOIN join_a\n      ON foo.id = join_a.id AND\n      join_a.id IN (\n        SELECT id\n        FROM sub_a\n        INNER JOIN sub_b\n          ON sub_a.id = sub_b.id\n            AND sub_b.id IN (\n              SELECT id\n              FROM sub_c\n              INNER JOIN sub_d ON sub_c.id IN (SELECT id from sub_e)\n            )\n      )\n    INNER JOIN join_b\n      ON foo.id = join_b.id AND\n      join_b.id IN (\n        SELECT id FROM sub_f\n      )\n",
			Tables:    &[]string{"join_b", "foo", "join_a", "sub_f", "sub_a", "sub_b", "sub_c", "sub_d", "sub_e"},
			Functions: &[]string{},
			FuncArgs:  &[]string{},
			Aliases:   &map[string]string{},
			CteNames:  &[]string{}},
		{
			Query:     "    SELECT *\n    FROM foo\n    INNER JOIN join_a\n      ON foo.id = join_a.id AND\n      join_a.id IN (\n        SELECT id\n        FROM sub_a\n        INNER JOIN sub_b\n          ON sub_a.id = sub_b.id\n            AND sub_b.id IN (\n              SELECT id\n              FROM sub_c\n              INNER JOIN sub_d ON sub_c.id IN (SELECT id from sub_e)\n            )\n      )\n    INNER JOIN join_b\n      ON foo.id = join_b.id AND\n      join_b.id IN (\n        SELECT id FROM sub_f\n      )\n",
			Tables:    &[]string{"join_b", "foo", "join_a", "sub_f", "sub_a", "sub_b", "sub_c", "sub_d", "sub_e"},
			Functions: &[]string{},
			FuncArgs:  &[]string{},
			Aliases:   &map[string]string{},
			CteNames:  &[]string{}},
		{
			Query:     "    with cte_a as (\n      select * from table_a\n    ), cte_b as (\n      select * from table_b\n    )\n\n    select id from table_c\n    left join cte_b on\n      table_c.id = cte_b.c_id\n    union\n    select * from cte_a\n",
			Tables:    &[]string{"table_c", "table_a", "table_b"},
			Functions: &[]string{},
			FuncArgs:  &[]string{},
			Aliases:   &map[string]string{},
			CteNames:  &[]string{"cte_a", "cte_b"}},
		{
			Query:     "    with cte_a as (\n      select * from table_a\n    ), cte_b as (\n      select * from table_b\n    )\n\n    select id from table_c\n    left join cte_b on\n      table_c.id = cte_b.c_id\n    except\n    select * from cte_a\n",
			Tables:    &[]string{"table_c", "table_a", "table_b"},
			Functions: &[]string{},
			FuncArgs:  &[]string{},
			Aliases:   &map[string]string{},
			CteNames:  &[]string{"cte_a", "cte_b"}},
		{
			Query:     "    with cte_a as (\n      select * from table_a\n    ), cte_b as (\n      select * from table_b\n    )\n\n    select id from table_c\n    left join cte_b on\n      table_c.id = cte_b.c_id\n    intersect\n    select * from cte_a\n",
			Tables:    &[]string{"table_c", "table_a", "table_b"},
			Functions: &[]string{},
			FuncArgs:  &[]string{},
			Aliases:   &map[string]string{},
			CteNames:  &[]string{"cte_a", "cte_b"}},
		{
			Query:     "    SELECT GREATEST(\n             date_trunc($1, $2::timestamptz) + $3::interval,\n             COALESCE(\n               (\n                 SELECT first_aggregate_starts_at\n                   FROM schema_aggregate_infos\n                  WHERE base_table = $4 LIMIT $5\n               ),\n               now() + $6::interval\n             )\n          ) AS first_hourly_start_ts\n",
			Tables:    &[]string{"schema_aggregate_infos"},
			Functions: &[]string{"date_trunc", "now"},
			FuncArgs:  &[]string{},
			Aliases:   &map[string]string{},
			CteNames:  &[]string{}},
		{
			Query:     "    SELECT\n      CASE\n        WHEN id IN (SELECT foo_id FROM when_a) THEN (SELECT MAX(id) FROM then_a)\n        WHEN id IN (SELECT foo_id FROM when_b) THEN (SELECT MAX(id) FROM then_b)\n        ELSE (SELECT MAX(id) FROM elsey)\n      END\n    FROM foo\n",
			Tables:    &[]string{"foo", "when_a", "when_b", "then_a", "then_b", "elsey"},
			Functions: &[]string{"max"},
			FuncArgs:  &[]string{"id"},
			Aliases:   &map[string]string{},
			CteNames:  &[]string{}},
		{
			Query:     "    SELECT 1\n    FROM   foo\n    WHERE  x = any(cast(array(SELECT a FROM bar) as bigint[]))\n        OR x = any(array(SELECT a FROM baz)::bigint[])\n",
			Tables:    &[]string{"foo", "bar", "baz"},
			Functions: &[]string{},
			FuncArgs:  &[]string{},
			Aliases:   &map[string]string{},
			CteNames:  &[]string{}},
		{
			Query:     "    SELECT *\n      FROM my_custom_func()\n",
			Tables:    &[]string{},
			Functions: &[]string{"my_custom_func"},
			FuncArgs:  &[]string{},
			Aliases:   &map[string]string{},
			CteNames:  &[]string{}},
		{
			Query:     "    SELECT *\n      FROM unnest($1::text[]) AS a(x)\n      LEFT OUTER JOIN LATERAL (\n        SELECT json_build_object($2, \"z\".\"z\")\n          FROM (\n            SELECT *\n              FROM (\n                SELECT row_to_json(\n                    (SELECT * FROM (SELECT public.my_function(b) FROM public.c) d)\n                )\n              ) e\n        ) f\n      ) AS g ON (1)\n",
			Tables:    &[]string{"public.c"},
			Functions: &[]string{"unnest", "json_build_object", "row_to_json", "public.my_function"},
			FuncArgs:  &[]string{},
			Aliases:   &map[string]string{},
			CteNames:  &[]string{}},
		{
			Query:     "DROP TYPE IF EXISTS repack.pk_something",
			Tables:    &[]string{},
			Functions: &[]string{},
			FuncArgs:  &[]string{},
			Aliases:   &map[string]string{},
			CteNames:  &[]string{}},
		{
			Query:     "COPY (SELECT test FROM abc) TO STDOUT WITH (FORMAT 'csv')",
			Tables:    &[]string{"abc"},
			Functions: &[]string{},
			FuncArgs:  &[]string{},
			Aliases:   &map[string]string{},
			CteNames:  &[]string{}},
		{
			Query:     "    insert into users(pk, name) values (1, 'bob');\n",
			Tables:    &[]string{"users"},
			Functions: &[]string{},
			FuncArgs:  &[]string{},
			Aliases:   &map[string]string{},
			CteNames:  &[]string{}},
		{
			Query:     "    insert into users(pk, name) select pk, name from other_users;\n",
			Tables:    &[]string{"users", "other_users"},
			Functions: &[]string{},
			FuncArgs:  &[]string{},
			Aliases:   &map[string]string{},
			CteNames:  &[]string{}},
		{
			Query:     "    with cte as (\n      select pk, name from other_users\n    )\n    insert into users(pk, name) select * from cte;\n",
			Tables:    &[]string{"users", "other_users"},
			Functions: &[]string{},
			FuncArgs:  &[]string{},
			Aliases:   &map[string]string{},
			CteNames:  &[]string{"cte"}},
		{
			Query:     "    update users set name = 'bob';\n",
			Tables:    &[]string{"users"},
			Functions: &[]string{},
			FuncArgs:  &[]string{},
			Aliases:   &map[string]string{},
			CteNames:  &[]string{}},
		{
			Query:     "    update users set name = (select name from other_users limit 1);\n",
			Tables:    &[]string{"users", "other_users"},
			Functions: &[]string{},
			FuncArgs:  &[]string{},
			Aliases:   &map[string]string{},
			CteNames:  &[]string{}},
		{
			Query:     "    with cte as (\n      select name from other_users limit 1\n    )\n    update users set name = (select name from cte);\n",
			Tables:    &[]string{"users", "other_users"},
			Functions: &[]string{},
			FuncArgs:  &[]string{},
			Aliases:   &map[string]string{},
			CteNames:  &[]string{"cte"}},
		{
			Query:     "    UPDATE users SET name = users_new.name\n    FROM users_new\n    INNER JOIN join_table ON join_table.user_id = new_users.id\n    WHERE users.id = users_new.id\n",
			Tables:    &[]string{"users", "users_new", "join_table"},
			Functions: &[]string{},
			FuncArgs:  &[]string{},
			Aliases:   &map[string]string{},
			CteNames:  &[]string{}},
		{
			Query:     "    DELETE FROM users;\n",
			Tables:    &[]string{"users"},
			Functions: &[]string{},
			FuncArgs:  &[]string{},
			Aliases:   &map[string]string{},
			CteNames:  &[]string{}},
		{
			Query:     "    DELETE FROM users USING foo\n      WHERE foo_id = foo.id AND foo.action = 'delete';\n",
			Tables:    &[]string{"users", "foo"},
			Functions: &[]string{},
			FuncArgs:  &[]string{},
			Aliases:   &map[string]string{},
			CteNames:  &[]string{}},
		{
			Query:     "    DELETE FROM users\n      WHERE foo_id IN (SELECT id FROM foo WHERE action = 'delete');\n",
			Tables:    &[]string{"users", "foo"},
			Functions: &[]string{},
			FuncArgs:  &[]string{},
			Aliases:   &map[string]string{},
			CteNames:  &[]string{}},
		{
			Query:     "    CREATE TABLE foo AS\n      SELECT * FROM bar;\n",
			Tables:    &[]string{"foo", "bar"},
			Functions: &[]string{},
			FuncArgs:  &[]string{},
			Aliases:   &map[string]string{},
			CteNames:  &[]string{}},
		{
			Query:     "    CREATE TABLE foo AS\n      SELECT id FROM bar UNION SELECT id from baz;\n",
			Tables:    &[]string{"foo", "bar", "baz"},
			Functions: &[]string{},
			FuncArgs:  &[]string{},
			Aliases:   &map[string]string{},
			CteNames:  &[]string{}},
		{
			Query:     "    CREATE TABLE foo AS\n      SELECT id FROM bar EXCEPT SELECT id from baz;\n",
			Tables:    &[]string{"foo", "bar", "baz"},
			Functions: &[]string{},
			FuncArgs:  &[]string{},
			Aliases:   &map[string]string{},
			CteNames:  &[]string{}},
		{
			Query:     "    CREATE TABLE foo AS\n      SELECT id FROM bar INTERSECT SELECT id from baz;\n",
			Tables:    &[]string{"foo", "bar", "baz"},
			Functions: &[]string{},
			FuncArgs:  &[]string{},
			Aliases:   &map[string]string{},
			CteNames:  &[]string{}},
		{
			Query:     "    PREPARE qux AS SELECT bar from foo\n",
			Tables:    &[]string{"foo"},
			Functions: &[]string{},
			FuncArgs:  &[]string{},
			Aliases:   &map[string]string{},
			CteNames:  &[]string{}},
		{
			Query:     "    CREATE TEMP TABLE foo AS SELECT 1;\n",
			Tables:    &[]string{"foo"},
			Functions: &[]string{},
			FuncArgs:  &[]string{},
			Aliases:   &map[string]string{},
			CteNames:  &[]string{}},
	}
	for i, input := range expectations {
		pr, err := pg_query.Parse(input.Query)
		if err != nil {
			if fmt.Sprint(err) != fmt.Sprint(input.Err) {
				t.Errorf("Parse() failed: input: [%d]%s, err %+v", i, input.Query, err)
				break
			}
			continue
		}
		tpr := pg_query.NewTaggedParseResult(pr)
		err = tpr.LoadObjects()
		if err != nil {
			t.Errorf("LoadObjects() err failed: input: [%d]%s, expected %+v, actual %+v", i, input.Query, input.Err, err)
			break
		}

		if input.Tables != nil {
			tables := tpr.GetTables()
			if len(tables) != len(*input.Tables) {
				t.Errorf("LoadObjects() tables failed: input: [%d]%s, expected %+v, actual %+v", i, input.Query, *input.Tables, tables)
				break
			}
			for j, table := range tables {
				if table.Name != (*input.Tables)[j] {
					t.Errorf("LoadObjects() tables failed: input: [%d]%s, expected %+v, actual %+v", i, input.Query, *input.Tables, tables)
					break
				}
			}
		}
		if input.Functions != nil {
			functions := tpr.GetFunctions()
			if len(functions) != len(*input.Functions) {
				t.Errorf("LoadObjects() functions failed: input: [%d]%s, expected %+v, actual %+v", i, input.Query, *input.Functions, functions)
				break
			}
			ok := true
			for j, function := range functions {
				if function.Name != (*input.Functions)[j] {
					ok = false
					break
				}
			}
			if !ok {
				t.Errorf("LoadObjects() functions failed: input: [%d]%s, expected %+v, actual %+v", i, input.Query, *input.Functions, functions)
				break
			}
		}
		if input.FuncArgs != nil {
			funcArgs := tpr.GetArgsFunctions()
			if len(funcArgs) != len(*input.FuncArgs) {
				t.Errorf("LoadObjects() funcArgs failed: input: [%d]%s, expected %+v, actual %+v", i, input.Query, *input.FuncArgs, funcArgs)
				break
			}
			ok := true
			for j, funcArg := range funcArgs {
				if funcArg.Name != (*input.FuncArgs)[j] {
					ok = false
					break
				}
			}
			if !ok {
				t.Errorf("LoadObjects() funcArgs failed: input: [%d]%s, expected %+v, actual %+v", i, input.Query, *input.FuncArgs, funcArgs)
				break
			}
		}
		if input.Aliases != nil {
			aliases := tpr.Aliases
			if len(aliases) != len(*input.Aliases) {
				t.Errorf("LoadObjects() aliases failed: input: [%d]%s, expected %+v, actual %+v", i, input.Query, *input.Aliases, aliases)
				break
			}
			for k, v := range aliases {
				if (*input.Aliases)[k] != v {
					t.Errorf("LoadObjects() aliases failed: input: [%d]%s, expected %+v, actual %+v", i, input.Query, *input.Aliases, aliases)
					break
				}
			}
		}
		if input.CteNames != nil {
			cteNames := tpr.CteNames
			if len(cteNames) != len(*input.CteNames) {
				t.Errorf("LoadObjects() cteNames failed: input: [%d]%s, expected %+v, actual %+v", i, input.Query, *input.CteNames, cteNames)
				break
			}
			for _, k := range *input.CteNames {
				if _, ok := cteNames[k]; !ok {
					t.Errorf("LoadObjects() cteNames failed: input: [%d]%s, expected %+v, actual %+v", i, input.Query, *input.CteNames, cteNames)
					break
				}
			}
		}

	}
}
