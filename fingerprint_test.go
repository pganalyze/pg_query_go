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
		"d64b6e1a3dad7d86c19ca57621fb29e3",
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
