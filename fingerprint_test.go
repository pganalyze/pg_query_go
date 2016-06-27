package pg_query_test

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"reflect"
	"testing"

	"github.com/lfittl/pg_query_go"
	nodes "github.com/lfittl/pg_query_go/nodes"
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

		actualTree, err := pg_query.Parse(test.Input)
		if err != nil {
			t.Errorf("Fingerprint(%s)\nparse error %s\n\n", test.Input, err)
		}

		ctx := nodes.NewFingerprintSubContext()
		for _, node := range actualTree.Statements {
			node.Fingerprint(ctx, nil, "")
		}
		if !reflect.DeepEqual(ctx.Sum(), test.ExpectedParts) {
			t.Errorf("Fingerprint(%s)\nexpected parts %#v\nactual parts %#v\n\n", test.Input, test.ExpectedParts, ctx.Sum())
		}

		actual := actualTree.Fingerprint()

		if string(actual) != test.ExpectedHash {
			t.Errorf("Fingerprint(%s)\nexpected %s\nactual %s\n\n", test.Input, test.ExpectedHash, actual)
		}
	}

	fmt.Printf("\n")
}
