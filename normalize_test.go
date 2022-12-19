package pg_query_test

import (
	"errors"
	"reflect"
	"testing"

	pg_query "github.com/pganalyze/pg_query_go/v2"
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
		errors.New("syntax error at or near \"$\""),
	},
}

func TestNormalizeError(t *testing.T) {
	for _, test := range normalizeErrorTests {
		_, actualErr := pg_query.Normalize(test.input)

		if actualErr == nil {
			t.Errorf("Normalize(%s)\nexpected error but none returned\n\n", test.input)
		} else if !reflect.DeepEqual(actualErr, test.expectedErr) {
			t.Errorf("Normalize(%s)\nexpected error %s\nactual error %s\n\n", test.input, test.expectedErr, actualErr)
		}
	}
}
