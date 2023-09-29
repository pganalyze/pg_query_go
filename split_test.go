package pg_query_test

import (
	"testing"

	pg_query "github.com/pganalyze/pg_query_go/v4"
)

var splitTests = []struct {
	name     string
	input    string
	expected []string
}{
	{
		name:  "basic split",
		input: "select * from a;select * from b;",
		expected: []string{
			"select * from a",
			"select * from b",
		},
	},
	{
		name:  "procedure",
		input: splitTestInput,
		expected: []string{
			splitExpected1,
			splitExpected2,
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
			actuals, err := pg_query.Split(test.input)
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
