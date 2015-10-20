package pg_query_test

import (
  "fmt"
  "encoding/json"
  "testing"
  "github.com/lfittl/pg_query.go"
  nodes "github.com/lfittl/pg_query.go/nodes"
)

var queryTests = []struct {
  input    string
  expected string
}{
  {"SELECT 1", `[{"SELECT": {"distinctClause": null, "intoClause": null, "targetList": [{"RESTARGET": {"name": null, "indirection": null, "val": {"A_CONST": {"val": 1, "location": 7}}, "location": 7}}], "fromClause": null, "whereClause": null, "groupClause": null, "havingClause": null, "windowClause": null, "valuesLists": null, "sortClause": null, "limitOffset": null, "limitCount": null, "lockingClause": null, "withClause": null, "op": 0, "all": false, "larg": null, "rarg": null}}]`},
}

func TestParse(t *testing.T) {
  fmt.Printf("%v\n", nodes.JsonKeyToNodeTag("SELECT"));

  stmt := nodes.SelectStmt{
    TargetList: []nodes.Node {
      nodes.ResTarget{
        Val: nodes.A_Const{
          Val: nodes.Value{
            Ival: 1,
          },
          Location: 7,
        },
        Location: 7,
      },
    },
  }
  fmt.Printf("%v\n", stmt)
  str, _ := json.Marshal(stmt)
  fmt.Printf("%s\n", str)
  fmt.Printf(pg_query.Parse("SELECT 1"))
  // [{"SELECT": {"distinctClause": null, "intoClause": null, "targetList": null, "fromClause": null, "whereClause": null, "groupClause": null, "havingClause": null, "windowClause": null, "valuesLists": null, "sortClause": null, "limitOffset": null, "limitCount": null, "lockingClause": null, "withClause": null, "op": 0, "all": false, "larg": null, "rarg": null}}]

  for _, test := range queryTests {
    actual := pg_query.Parse(test.input)
    if actual != test.expected {
      t.Errorf("Parse(%s): expected %s, actual %s", test.input, test.expected, actual)
    }
  }
}
