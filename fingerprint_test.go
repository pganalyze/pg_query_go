//go:build cgo
// +build cgo

package pg_query_test

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"testing"

	pg_query "github.com/pganalyze/pg_query_go/v6"
)

// trailingJSONCComment matches a JSON5/JSONC-style "// ..." comment
// immediately following a comma at the end of a line — the shape libpg_query
// 18.0.0's own testdata/fingerprint.json now ships (e.g. `"disableOnMsvc":
// true, // Doesn't work because of C2026: string too big`), which Go's
// strict encoding/json rejects outright. Anchored on a preceding comma
// specifically so a SQL fixture string that happens to contain "//" (e.g. a
// URL literal) is never misidentified: such a string is followed by a
// closing quote and comma, not by more unquoted text ending the line.
var trailingJSONCComment = regexp.MustCompile(`,(\s*)//[^\n]*`)

func stripTrailingJSONCComments(data []byte) []byte {
	return trailingJSONCComment.ReplaceAll(data, []byte(",$1"))
}

type fingerprintTest struct {
	Input         string
	ExpectedParts []string
	ExpectedHash  string
}

// TestFingerprint has 6 known-failing cases as of the PG18 (libpg_query
// 18.0.0) vendoring update: 3 alias-invariance cases, CREATE TEMPORARY
// TABLE ... ON COMMIT DROP, and 2 MERGE statements — all upstream's own
// golden hashes in testdata/fingerprint.json, not something this fork
// computes. Deliberately left red rather than silently adjusted or
// skipped: nothing in pg_query_go/v6's Fingerprint or FingerprintToUInt64
// is exercised by github.com/dullkingsman/dpg (confirmed via a full-repo
// grep before accepting this gap), so there was no way to determine
// "what the correct hash actually is" independent of trusting upstream's
// own possibly-stale fixture — worth another look if/when this fork ever
// picks up an upstream v7 release that fixes it, or if a consumer starts
// relying on Fingerprint.
func TestFingerprint(t *testing.T) {
	var fingerprintTests []fingerprintTest

	file, err := ioutil.ReadFile("./testdata/fingerprint.json")
	if err != nil {
		t.Errorf("Could not load test file: %v\n", err)
	}
	file = stripTrailingJSONCComments(file)

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
