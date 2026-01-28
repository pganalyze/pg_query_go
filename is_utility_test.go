//go:build cgo
// +build cgo

package pg_query_test

import (
	"reflect"
	"testing"

	pg_query "github.com/pganalyze/pg_query_go/v6"
	"github.com/pganalyze/pg_query_go/v6/parser"
)

var isUtilityStmtTests = []struct {
	input    string
	expected []bool
}{
	// DML statements (not utility)
	{
		"SELECT 1",
		[]bool{false},
	},
	{
		"INSERT INTO t (a) VALUES (1)",
		[]bool{false},
	},
	{
		"UPDATE t SET a = 1",
		[]bool{false},
	},
	{
		"DELETE FROM t",
		[]bool{false},
	},
	// Utility statements
	{
		"SHOW fsync",
		[]bool{true},
	},
	{
		"SET fsync = off",
		[]bool{true},
	},
	{
		"CREATE TABLE t (a int)",
		[]bool{true},
	},
	{
		"DROP TABLE t",
		[]bool{true},
	},
	// Multi-statement input
	{
		"SELECT 1; SELECT 2",
		[]bool{false, false},
	},
	{
		"SELECT 1; SHOW fsync",
		[]bool{false, true},
	},
	{
		"SHOW fsync; SELECT 1",
		[]bool{true, false},
	},
	{
		"SET a = 1; SET b = 2",
		[]bool{true, true},
	},
}

func TestIsUtilityStmt(t *testing.T) {
	for _, test := range isUtilityStmtTests {
		actual, err := pg_query.IsUtilityStmt(test.input)

		if err != nil {
			t.Errorf("IsUtilityStmt(%s)\nerror %s\n\n", test.input, err)
		} else if !reflect.DeepEqual(actual, test.expected) {
			t.Errorf("IsUtilityStmt(%s)\nexpected %v\nactual %v\n\n", test.input, test.expected, actual)
		}
	}
}

var isUtilityStmtErrorTests = []struct {
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
}

func TestIsUtilityStmtError(t *testing.T) {
	for _, test := range isUtilityStmtErrorTests {
		_, actualErr := pg_query.IsUtilityStmt(test.input)

		if actualErr == nil {
			t.Errorf("IsUtilityStmt(%s)\nexpected error but none returned\n\n", test.input)
		} else {
			exp := test.expectedErr.(*parser.Error)
			act := actualErr.(*parser.Error)
			act.Lineno = 0 // Line number is architecture dependent, so we ignore it
			if !reflect.DeepEqual(act, exp) {
				t.Errorf(
					"IsUtilityStmt(%s)\nexpected error %s at %d (%s:%d), func: %s, context: %s\nactual error %+v at %d (%s:%d), func: %s, context: %s\n\n",
					test.input,
					exp.Message, exp.Cursorpos, exp.Filename, exp.Lineno, exp.Funcname, exp.Context,
					act.Message, act.Cursorpos, act.Filename, act.Lineno, act.Funcname, act.Context)
			}
		}
	}
}
