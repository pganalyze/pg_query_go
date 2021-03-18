package pg_query_test

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"testing"

	pg_query "github.com/pganalyze/pg_query_go/v2"
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
	}

	fmt.Printf("\n")
}
