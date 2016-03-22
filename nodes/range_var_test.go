package pg_query_test

import (
	"encoding/json"
	"reflect"
	"testing"

	nodes "github.com/lfittl/pg_query_go/nodes"
)

func strPtr(str string) *string {
	return &str
}

var rangeVarTests = []struct {
	jsonText     string
	expectedNode nodes.RangeVar
}{
	{
		`{"schemaname": null, "relname": "x", "inhOpt": 2, "relpersistence": "p", "alias": null, "location": 14}`,
		nodes.RangeVar{
			Relname:        strPtr("x"),
			InhOpt:         nodes.INH_DEFAULT,
			Relpersistence: 'p',
			Location:       14,
		},
	},
}

func TestRangeVar(t *testing.T) {
	for _, test := range rangeVarTests {
		var actualTree nodes.RangeVar
		err := json.Unmarshal([]byte(test.jsonText), &actualTree)

		if err != nil {
			t.Errorf("Unmarshal(%s)\nerror %s\n\n", test.jsonText, err)
		} else if !reflect.DeepEqual(actualTree, test.expectedNode) {
			t.Errorf("Unmarshal(%s)\nexpected %s\nactual %s\n\n", test.jsonText, test.expectedNode, actualTree)
		}
	}
}
