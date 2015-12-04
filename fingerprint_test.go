package pg_query_test

import (
	"testing"

	"github.com/lfittl/pg_query.go"
)

var fingerprintTests = []struct {
	input    string
	expected string
}{
	{
		"SELECT 1",
		"4a76edca1a5766d542e5bde019dc8a7ee4f51726",
	},
	{
		"SELECT 2",
		"4a76edca1a5766d542e5bde019dc8a7ee4f51726",
	},
	{
		"SELECT COUNT(DISTINCT id), * FROM targets WHERE something IS NOT NULL AND elsewhere::interval < now()",
		"feb7587c16f46a5fd771c841cf8cb66aa21c692a",
	},
}

func TestFingerprint(t *testing.T) {
	for _, test := range fingerprintTests {
		actualTree, err := pg_query.Parse(test.input)
		if err != nil {
			t.Errorf("Fingerprint(%s)\nparse error %s\n\n", test.input, err)
		}

		actual := actualTree.Fingerprint()

		if string(actual) != test.expected {
			t.Errorf("Fingerprint(%s)\nexpected %s\nactual %s\n\n", test.input, test.expected, actual)
		}
	}
}
