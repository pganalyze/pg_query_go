package pg_query_test

import (
	"reflect"
	"testing"

	"github.com/lfittl/pg_query.go"
	nodes "github.com/lfittl/pg_query.go/nodes"
)

var fingerprintTests = []struct {
	input         string
	expectedParts []string
	expectedHash  string
}{
	{
		"SELECT 1",
		[]string{"SelectStmt", "false", "0", "ResTarget"},
		"31dc5500dc27777a26160cb1b0faa11495f150d8",
	},
	{
		"SELECT 2",
		[]string{"SelectStmt", "false", "0", "ResTarget"},
		"31dc5500dc27777a26160cb1b0faa11495f150d8",
	},
	{
		"SELECT COUNT(DISTINCT id), * FROM targets WHERE something IS NOT NULL AND elsewhere::interval < now()",
		[]string{"SelectStmt", "false", "RangeVar", "2", "targets", "p", "0", "ResTarget", "ColumnRef",
			"A_Star", "ResTarget", "FuncCall", "true", "false", "false", "ColumnRef", "String",
			"id", "false", "String", "count", "A_Expr", "1", "NullTest", "ColumnRef", "String",
			"something", "false", "1", "A_Expr", "0", "TypeCast", "ColumnRef", "String", "elsewhere",
			"TypeName", "String", "pg_catalog", "String", "interval", "false", "false", "0", "-1",
			"String", "<", "FuncCall", "false", "false", "false", "false", "String", "now"},
		"1e014ccea580bb5dea8b4a66893b3c508d6261f0",
	},
	{
		"INSERT INTO test (a, b) VALUES (?, ?)",
		[]string{"InsertStmt", "ResTarget", "a", "ResTarget", "b", "RangeVar", "2", "test", "p", "SelectStmt", "false", "0"},
		"f3f2847a56d9b67f11e1905d2365bc627f852220",
	},
	{
		"INSERT INTO test (b, a) VALUES (?, ?)",
		[]string{"InsertStmt", "ResTarget", "a", "ResTarget", "b", "RangeVar", "2", "test", "p", "SelectStmt", "false", "0"},
		"f3f2847a56d9b67f11e1905d2365bc627f852220",
	},
	{
		"SELECT b AS x, a AS y FROM z",
		[]string{"SelectStmt", "false", "RangeVar", "2", "z", "p", "0", "ResTarget", "ColumnRef", "String", "a", "ResTarget", "ColumnRef", "String", "b"},
		"7c361dd7a746418464fdf666cfae7be6a0f873aa",
	},
	{
		"SELECT * FROM x WHERE y IN (?)",
		[]string{"SelectStmt", "false", "RangeVar", "2", "x", "p", "0", "ResTarget", "ColumnRef", "A_Star", "A_Expr", "9", "ColumnRef", "String", "y", "String", "="},
		"d15431e54000f340c2bfca70ed3a0f31b2e55061",
	},
	{
		"SELECT * FROM x WHERE y IN (?, ?, ?)",
		[]string{"SelectStmt", "false", "RangeVar", "2", "x", "p", "0", "ResTarget", "ColumnRef", "A_Star", "A_Expr", "9", "ColumnRef", "String", "y", "String", "="},
		"d15431e54000f340c2bfca70ed3a0f31b2e55061",
	},
	{
		"SELECT * FROM x WHERE y IN ( ?::uuid )",
		[]string{"SelectStmt", "false", "RangeVar", "2", "x", "p", "0", "ResTarget", "ColumnRef", "A_Star", "A_Expr", "9", "ColumnRef", "String", "y", "String", "=", "TypeCast", "TypeName", "String", "uuid", "false", "false", "0", "-1"},
		"bd7dcab89d5a8ad04b5f7e352030f47d5abd1eab",
	},
	{
		"SELECT * FROM x WHERE y IN ( ?::uuid, ?::uuid, ?::uuid )",
		[]string{"SelectStmt", "false", "RangeVar", "2", "x", "p", "0", "ResTarget", "ColumnRef", "A_Star", "A_Expr", "9", "ColumnRef", "String", "y", "String", "=", "TypeCast", "TypeName", "String", "uuid", "false", "false", "0", "-1"},
		"bd7dcab89d5a8ad04b5f7e352030f47d5abd1eab",
	},
}

func TestFingerprint(t *testing.T) {
	for _, test := range fingerprintTests {
		actualTree, err := pg_query.Parse(test.input)
		if err != nil {
			t.Errorf("Fingerprint(%s)\nparse error %s\n\n", test.input, err)
		}

		ctx := &nodes.FingerprintSubContext{}
		for _, node := range actualTree.Statements {
			node.Fingerprint(ctx, "")
		}
		if !reflect.DeepEqual(ctx.Sum(), test.expectedParts) {
			t.Errorf("Fingerprint(%s)\nexpected parts %v\nactual parts %v\n\n", test.input, test.expectedParts, ctx.Sum())
		}

		actual := actualTree.Fingerprint()

		if string(actual) != test.expectedHash {
			t.Errorf("Fingerprint(%s)\nexpected %s\nactual %s\n\n", test.input, test.expectedHash, actual)
		}
	}
}
