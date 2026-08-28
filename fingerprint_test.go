//go:build cgo
// +build cgo

package pg_query_test

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"
	"testing"

	pg_query "github.com/pganalyze/pg_query_go/v6"
	"github.com/tailscale/hujson"
)

type fingerprintTest struct {
	Input         string
	ExpectedParts []string
	ExpectedHash  string
}

func TestFingerprint(t *testing.T) {
	var fingerprintTests []fingerprintTest

	file, err := ioutil.ReadFile("./testdata/fingerprint.json")
	if err != nil {
		t.Errorf("Could not load test file: %v\n", err)
	}

	// The upstream test file may contain comments, which encoding/json rejects
	file, err = hujson.Standardize(file)
	if err != nil {
		t.Errorf("Could not parse test file: %v\n", err)
	}

	err = json.Unmarshal(file, &fingerprintTests)
	if err != nil {
		t.Errorf("Could not parse test file: %v\n", err)
	}

	for _, test := range fingerprintTests {
		fmt.Printf(".")

		fingerprint, err := pg_query.Fingerprint(test.Input)
		if err != nil {
			t.Errorf("Fingerprint(%s)\nparse error %s\n\n", test.Input, err)
		}

		if string(fingerprint) != test.ExpectedHash {
			t.Errorf("Fingerprint(%s)\nexpected %s\nactual %s\n\n", test.Input, test.ExpectedHash, fingerprint)
		}

		fingerprintInt, err := pg_query.FingerprintToUInt64(test.Input)
		if err != nil {
			t.Errorf("FingerprintToUInt64(%s)\nparse error %s\n\n", test.Input, err)
		}

		expectedInt, _ := strconv.ParseUint(test.ExpectedHash, 16, 64)

		if fingerprintInt != expectedInt {
			t.Errorf("FingerprintToUInt64(%s)\nexpected %d\nactual %d\n\n", test.Input, expectedInt, fingerprintInt)
		}
	}

	fmt.Printf("\n")
}

// Expected values match the upstream libpg_query fingerprint option tests
// (test/fingerprint_opts_tests.c)
var fingerprintOptsTests = []struct {
	input    string
	opts     pg_query.FingerprintOption
	expected string
}{
	// By default, 2+ consecutive digits in the relation name are ignored (these two match)
	{"SELECT * FROM orders_2024_01", pg_query.FingerprintDefault, "0e612f391ad711b8"},
	{"SELECT * FROM orders_2024_02", pg_query.FingerprintDefault, "0e612f391ad711b8"},
	// With FingerprintRelnameFull the full relation name is fingerprinted (these two differ)
	{"SELECT * FROM orders_2024_01", pg_query.FingerprintRelnameFull, "3cc2d1ca3f22c9bf"},
	{"SELECT * FROM orders_2024_02", pg_query.FingerprintRelnameFull, "291f96ac98cf4c38"},
	// By default (Postgres 18+ behavior), the alias replaces the relation name, and the schema name is ignored
	{"SELECT * FROM sales", pg_query.FingerprintDefault, "4d93c901b91cb364"},
	{"SELECT * FROM public.sales", pg_query.FingerprintDefault, "4d93c901b91cb364"},
	{"SELECT * FROM sales s", pg_query.FingerprintDefault, "93a3bbe18171c380"},
	// With FingerprintRangeVarIgnoreAliases, aliases are ignored (matches "SELECT * FROM sales" above)
	{"SELECT * FROM sales s", pg_query.FingerprintRangeVarIgnoreAliases, "4d93c901b91cb364"},
	// With FingerprintRangeVarIncludeSchema, the schema name is fingerprinted (differs from "SELECT * FROM sales" above)
	{"SELECT * FROM public.sales", pg_query.FingerprintRangeVarIncludeSchema, "78d676e53f612747"},
	// ... whilst aliases still replace the relation name
	{"SELECT * FROM public.sales s", pg_query.FingerprintRangeVarIncludeSchema, "9cf22829ca3b350b"},
	// PG17_COMPAT matches the fingerprint from libpg_query 17
	{"SELECT * FROM x AS a, y AS b", pg_query.FingerprintRangeVarPG17Compat, "4e9acae841dae228"},
	// All flags combined
	{"SELECT * FROM public.orders_2024_01 o", pg_query.FingerprintRangeVarPG17Compat | pg_query.FingerprintRelnameFull, "115077f8a9c3c10d"},
}

func TestFingerprintWithOpts(t *testing.T) {
	for _, test := range fingerprintOptsTests {
		fingerprint, err := pg_query.FingerprintWithOpts(test.input, test.opts)
		if err != nil {
			t.Errorf("FingerprintWithOpts(%s, %d)\nparse error %s\n\n", test.input, test.opts, err)
		}

		if fingerprint != test.expected {
			t.Errorf("FingerprintWithOpts(%s, %d)\nexpected %s\nactual %s\n\n", test.input, test.opts, test.expected, fingerprint)
		}

		fingerprintInt, err := pg_query.FingerprintToUInt64WithOpts(test.input, test.opts)
		if err != nil {
			t.Errorf("FingerprintToUInt64WithOpts(%s, %d)\nparse error %s\n\n", test.input, test.opts, err)
		}

		expectedInt, _ := strconv.ParseUint(test.expected, 16, 64)

		if fingerprintInt != expectedInt {
			t.Errorf("FingerprintToUInt64WithOpts(%s, %d)\nexpected %d\nactual %d\n\n", test.input, test.opts, expectedInt, fingerprintInt)
		}
	}
}

var hashTests = []struct {
	input    string
	seed     uint64
	expected uint64
}{
	{
		"TEST",
		0,
		11717748491247689214,
	},
	{
		"TEST",
		42,
		10412276358662179996,
	},
	{
		"Something else",
		0,
		14679351602596009561,
	},
}

func TestHashXXH3_64(t *testing.T) {
	for _, test := range hashTests {
		actual := pg_query.HashXXH3_64([]byte(test.input), test.seed)

		if actual != test.expected {
			t.Errorf("HashXXH3_64(%s)\nexpected %d\nactual %d\n\n", test.input, test.expected, actual)
		}
	}
}
