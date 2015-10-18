package pg_query_test

import (
  "testing"
  "github.com/lfittl/pg_query.go"
)

var queryTests = []struct {
  input    string
  expected string
}{
  {"SELECT 1", `[{"SELECT": {"distinctClause": null, "intoClause": null, "targetList": [{"RESTARGET": {"name": null, "indirection": null, "val": {"A_CONST": {"val": 1, "location": 7}}, "location": 7}}], "fromClause": null, "whereClause": null, "groupClause": null, "havingClause": null, "windowClause": null, "valuesLists": null, "sortClause": null, "limitOffset": null, "limitCount": null, "lockingClause": null, "withClause": null, "op": 0, "all": false, "larg": null, "rarg": null}}]`},
}

func TestParse(t *testing.T) {
  for _, test := range queryTests {
    actual := pg_query.Parse(test.input)
    if actual != test.expected {
      t.Errorf("Parse(%s): expected %s, actual %s", test.input, test.expected, actual)
    }
  }
}
