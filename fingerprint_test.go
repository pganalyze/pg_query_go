//go:build cgo
// +build cgo

package pg_query_test

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"
	"testing"

	pg_query "github.com/xataio/pg_query_go/v6"
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
