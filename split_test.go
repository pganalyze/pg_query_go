//go:build cgo
// +build cgo

package pg_query_test

import (
	"testing"

	pg_query "github.com/xataio/pg_query_go/v6"
)

var splitTests = []struct {
	name      string
	splitFunc func(string, bool) ([]string, error)
	input     string
	trimSpace bool
	expected  []string
}{
	{
		name:      "splitWithParser - basic split",
		splitFunc: pg_query.SplitWithParser,
		input:     "select * from a;select * from b;",
		trimSpace: true,
		expected: []string{
			"select * from a",
			"select * from b",
		},
	},
	{
		name:      "splitWithParser - procedure",
		splitFunc: pg_query.SplitWithParser,
		input:     splitTestInput,
		trimSpace: true,
		expected: []string{
			splitExpected1,
			splitExpected2,
		},
	},
	{
		name:      "splitWithParser - basic split, no trim",
		splitFunc: pg_query.SplitWithParser,
		input:     "   select * from a;select * from b;",
		trimSpace: false,
		expected: []string{
			"   select * from a",
			"select * from b",
		},
	},
	{
		name:      "splitWithScanner - basic split",
		splitFunc: pg_query.SplitWithScanner,
		input:     "select * from a;select * from b;",
		trimSpace: true,
		expected: []string{
			"select * from a",
			"select * from b",
		},
	},
	{
		name:      "splitWithScanner - procedure",
		splitFunc: pg_query.SplitWithScanner,
		input:     splitTestInput,
		trimSpace: true,
		expected: []string{
			splitExpected1,
			splitExpected2,
		},
	},
	{
		name:      "splitWithScanner - basic split, no trim",
		splitFunc: pg_query.SplitWithScanner,
		input:     "   select * from a;select * from b;",
		trimSpace: false,
		expected: []string{
			"   select * from a",
			"select * from b",
		},
	},
}

var (
	splitTestInput = `UPDATE client SET name = "John Doe" WHERE id = 1;

CREATE OR REPLACE FUNCTION increment(i integer) RETURNS integer AS $$
	BEGIN
		RETURN i + 1;
    END;
$$ LANGUAGE plpgsql;
`
	splitExpected1 = `UPDATE client SET name = "John Doe" WHERE id = 1`
	splitExpected2 = `CREATE OR REPLACE FUNCTION increment(i integer) RETURNS integer AS $$
	BEGIN
		RETURN i + 1;
    END;
$$ LANGUAGE plpgsql`
)

func TestSplit(t *testing.T) {
	for _, test := range splitTests {
		t.Run(test.name, func(t *testing.T) {
			actuals, err := test.splitFunc(test.input, test.trimSpace)
			if err != nil {
				t.Error(err)
			}
			if len(actuals) != len(test.expected) {
				t.Error("unexpected number of results")
			}
			for i, actual := range actuals {
				if actual != test.expected[i] {
					t.Errorf("expected: [%s], actual: [%s]", test.expected[i], actual)
				}
			}
		})
	}
}
