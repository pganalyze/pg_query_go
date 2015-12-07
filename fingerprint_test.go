package pg_query_test

import (
	"reflect"
	"testing"

	"github.com/lfittl/pg_query.go"
)

var fingerprintTests = []struct {
	input         string
	expectedParts []string
	expectedHash  string
}{
	{
		"SELECT 1",
		[]string{"SELECT", "false", "0", "RESTARGET"},
		"4a76edca1a5766d542e5bde019dc8a7ee4f51726",
	},
	{
		"SELECT 2",
		[]string{"SELECT", "false", "0", "RESTARGET"},
		"4a76edca1a5766d542e5bde019dc8a7ee4f51726",
	},
	{
		"SELECT COUNT(DISTINCT id), * FROM targets WHERE something IS NOT NULL AND elsewhere::interval < now()",
		[]string{"SELECT", "false", "RANGEVAR", "2", "targets", "0", "RESTARGET", "FUNCCALL", "true", "false", "false", "COLUMNREF", "653", "id", "0", "false", "653", "count", "0", "RESTARGET", "COLUMNREF", "A_STAR"},
		"feb7587c16f46a5fd771c841cf8cb66aa21c692a",
	},
}

type FingerprintTestContext struct {
	parts []string
}

func (ctx *FingerprintTestContext) WriteString(str string) {
	ctx.parts = append(ctx.parts, str)
}

func TestFingerprint(t *testing.T) {
	for _, test := range fingerprintTests {
		actualTree, err := pg_query.Parse(test.input)
		if err != nil {
			t.Errorf("Fingerprint(%s)\nparse error %s\n\n", test.input, err)
		}

		ctx := &FingerprintTestContext{}
		for _, node := range actualTree.Statements {
			node.Fingerprint(ctx)
		}
		if !reflect.DeepEqual(ctx.parts, test.expectedParts) {
			t.Errorf("Fingerprint(%s)\nexpected parts %v\nactual parts %v\n\n", test.input, test.expectedParts, ctx.parts)
		}

		actual := actualTree.Fingerprint()

		if string(actual) != test.expectedHash {
			t.Errorf("Fingerprint(%s)\nexpected %s\nactual %s\n\n", test.input, test.expectedHash, actual)
		}
	}
}
