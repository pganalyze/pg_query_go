package pg_query_test

import (
	"reflect"
	"testing"

	pg_query "github.com/pganalyze/pg_query_go/v4"
	"github.com/pganalyze/pg_query_go/v4/parser"
)

var normalizeTests = []struct {
	input    string
	expected string
}{
	{
		"SELECT 1",
		"SELECT $1",
	},
}

func TestNormalize(t *testing.T) {
	for _, test := range normalizeTests {
		actual, err := pg_query.Normalize(test.input)

		if err != nil {
			t.Errorf("Normalize(%s)\nerror %s\n\n", test.input, err)
		} else if !reflect.DeepEqual(actual, test.expected) {
			t.Errorf("Normalize(%s)\nexpected %s\nactual %s\n\n", test.input, test.expected, actual)
		}
	}
}

var normalizeErrorTests = []struct {
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

func TestNormalizeError(t *testing.T) {
	for _, test := range normalizeErrorTests {
		_, actualErr := pg_query.Normalize(test.input)

		if actualErr == nil {
			t.Errorf("Normalize(%s)\nexpected error but none returned\n\n", test.input)
		} else {
			exp := test.expectedErr.(*parser.Error)
			act := actualErr.(*parser.Error)
			act.Lineno = 0 // Line number is architecture dependent, so we ignore it
			if !reflect.DeepEqual(act, exp) {
				t.Errorf(
					"Normalize(%s)\nexpected error %s at %d (%s:%d), func: %s, context: %s\nactual error %+v at %d (%s:%d), func: %s, context: %s\n\n",
					test.input,
					exp.Message, exp.Cursorpos, exp.Filename, exp.Lineno, exp.Funcname, exp.Context,
					act.Message, act.Cursorpos, act.Filename, act.Lineno, act.Funcname, act.Context)
			}
		}
	}
}
