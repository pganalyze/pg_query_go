//go:build cgo
// +build cgo

package pg_query_test

import (
	"testing"

	pg_query "github.com/pganalyze/pg_query_go/v6"
)

// These raw-parse-mode entry points (RAW_PARSE_PLPGSQL_EXPR and
// RAW_PARSE_PLPGSQL_ASSIGN1/2/3) are already used internally by
// ParsePlPgSqlToJSON to compile each PLpgSQL_expr node it finds, but were
// never exposed as standalone Go functions. They're useful on their own for
// re-parsing an individual PLpgSQL_expr fragment (the "query" text captured
// by ParsePlPgSqlToJSON) once its parseMode is known, e.g. to canonicalize
// or otherwise analyze it independently of the rest of the function body.

func TestParsePlPgSqlExpr(t *testing.T) {
	tree, err := pg_query.ParsePlPgSqlExpr("1 + 1")
	if err != nil {
		t.Fatalf("ParsePlPgSqlExpr: %s", err)
	}
	if len(tree.Stmts) != 1 {
		t.Fatalf("expected 1 statement, got %d", len(tree.Stmts))
	}
}

func TestParsePlPgSqlExprWhitespaceInvariant(t *testing.T) {
	a, err := pg_query.ParsePlPgSqlExpr("v_name || '/' || v_version")
	if err != nil {
		t.Fatalf("parse a: %s", err)
	}
	b, err := pg_query.ParsePlPgSqlExpr("v_name||'/'||v_version")
	if err != nil {
		t.Fatalf("parse b: %s", err)
	}

	outA, err := pg_query.Deparse(a)
	if err != nil {
		t.Fatalf("deparse a: %s", err)
	}
	outB, err := pg_query.Deparse(b)
	if err != nil {
		t.Fatalf("deparse b: %s", err)
	}
	if outA != outB {
		t.Errorf("expected whitespace-only variants to deparse identically, got %q vs %q", outA, outB)
	}
}

func TestParsePlPgSqlExprRejectsFullStatement(t *testing.T) {
	if _, err := pg_query.ParsePlPgSqlExpr("SELECT 1"); err == nil {
		t.Error("expected an error parsing a full statement as a bare expression")
	}
}

func TestParsePlPgSqlAssign1(t *testing.T) {
	tree, err := pg_query.ParsePlPgSqlAssign1("n := v_name || '/' || v_version")
	if err != nil {
		t.Fatalf("ParsePlPgSqlAssign1: %s", err)
	}
	if len(tree.Stmts) != 1 {
		t.Fatalf("expected 1 statement, got %d", len(tree.Stmts))
	}
	assign := tree.Stmts[0].Stmt.GetPlassignStmt()
	if assign == nil {
		t.Fatal("expected a PLAssignStmt")
	}
	if assign.Name != "n" {
		t.Errorf("expected target name %q, got %q", "n", assign.Name)
	}
	if assign.Val == nil {
		t.Fatal("expected Val to be populated")
	}

	out, err := pg_query.Deparse(&pg_query.ParseResult{
		Stmts: []*pg_query.RawStmt{
			{Stmt: &pg_query.Node{Node: &pg_query.Node_SelectStmt{SelectStmt: assign.Val}}},
		},
	})
	if err != nil {
		t.Fatalf("deparse Val: %s", err)
	}
	if out == "" {
		t.Error("expected a non-empty deparse of the assignment's right-hand side")
	}
}

func TestParsePlPgSqlAssign2DottedTarget(t *testing.T) {
	tree, err := pg_query.ParsePlPgSqlAssign2("rec.field := rec.field + 1")
	if err != nil {
		t.Fatalf("ParsePlPgSqlAssign2: %s", err)
	}
	assign := tree.Stmts[0].Stmt.GetPlassignStmt()
	if assign == nil {
		t.Fatal("expected a PLAssignStmt")
	}
	if assign.Name != "rec" {
		t.Errorf("expected target name %q, got %q", "rec", assign.Name)
	}
	if len(assign.Indirection) != 1 {
		t.Errorf("expected 1 indirection entry for a two-part target, got %d", len(assign.Indirection))
	}
}

func TestParsePlPgSqlAssign3(t *testing.T) {
	if _, err := pg_query.ParsePlPgSqlAssign3("a.b.c := 1"); err != nil {
		t.Fatalf("ParsePlPgSqlAssign3: %s", err)
	}
}
