package pg_query

import (
  "encoding/json"
  "reflect"
  "testing"
)

type TestStruct struct {
  Val Value	`json:"val"`
  Location int `json:"location"`
  Dummies []Node `json:"dummies"`
}

func (test *TestStruct) UnmarshalJSON(input []byte) error {
  return UnmarshalNodeFieldJSON(input, test)
}

func (test TestStruct) Deparse() string {
  panic("Not Implemented")
}

var testString = "test"

var queryTests = []struct {
  input    string
  expected TestStruct
}{
  {`{"location": 10}`, TestStruct{ Location: 10 }},
  {`{"val": 8}`, TestStruct{ Val: Value{ Type: T_Integer, Ival: 8 } }},
  {`{"dummies": [{"RESTARGET": {"name": "test"}}]}`, TestStruct{ Dummies: []Node{ResTarget{ Name: &testString } }}},
}

func TestUnmarshalNodeFieldJSON(t *testing.T) {
  for _, test := range queryTests {
    var actual TestStruct
    err := json.Unmarshal([]byte(test.input), &actual)
    if (err != nil) {
      t.Errorf("Unmarshal(%s)\nerror: %s\n\n", err)
    } else if (!reflect.DeepEqual(actual, test.expected)) {
      t.Errorf("Unmarshal(%s)\nexpected: %s\nactual: %s\n\n", test.input, test.expected, actual)
    }
  }
}
