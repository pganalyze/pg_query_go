package pg_query_test

import (
	"encoding/json"
	"reflect"
	"testing"

	nodes "github.com/lfittl/pg_query.go/nodes"
)

func strPtr(str string) *string {
	return &str
}

type TestStruct struct {
	Val        nodes.Value  `json:"val"`
	Location   int          `json:"location"`
	Dummies    []nodes.Node `json:"dummies"`
	MaybeStr   *string      `json:"maybe_str"`
	LittleByte byte         `json:"little_byte"`
}

func (test *TestStruct) UnmarshalJSON(input []byte) error {
	return nodes.UnmarshalNodeFieldJSON(input, test)
}

func (test TestStruct) Deparse() string {
	panic("Not Implemented")
}

var testString = "test"

var nodeTests = []struct {
	input    string
	expected TestStruct
}{
	{`{"location": 10}`, TestStruct{Location: 10}},
	{`{"val": 8}`, TestStruct{Val: nodes.Value{Type: nodes.T_Integer, Ival: 8}}},
	{`{"dummies": [{"RESTARGET": {"name": "test"}}]}`, TestStruct{Dummies: []nodes.Node{nodes.ResTarget{Name: &testString}}}},
	{`{"maybe_str": "x"}`, TestStruct{MaybeStr: strPtr("x")}},
	{`{"little_byte": "x"}`, TestStruct{LittleByte: 'x'}},
}

func TestNode(t *testing.T) {
	for _, test := range nodeTests {
		var actual TestStruct
		err := json.Unmarshal([]byte(test.input), &actual)
		if err != nil {
			t.Errorf("Unmarshal(%s)\nerror: %s\n\n", test.input, err)
		} else if !reflect.DeepEqual(actual, test.expected) {
			t.Errorf("Unmarshal(%s)\nexpected: %v\nactual: %v\n\n", test.input, test.expected, actual)
		}
	}
}
