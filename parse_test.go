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

  parsetree_list := pg_query.ParsetreeList{
    Statements: []nodes.Node{
      nodes.SelectStmt{
        TargetList: []nodes.Node{
          nodes.ResTarget{
            Val: nodes.A_Const{
              Val: nodes.Value{
                Type: nodes.T_Integer,
                Ival: 1,
              },
              Location: 7,
            },
            Location: 7,
          },
        },
      },
    },
  }


  fmt.Printf("MARSHAL\n")
  str, _ := json.Marshal(&parsetree_list)
  fmt.Printf("%s\n", str)
  fmt.Printf("%s\n", pg_query.Parse("SELECT 1"))

  fmt.Printf("UNMARSHAL\n")
  var new_parsetree_list pg_query.ParsetreeList
  err := json.Unmarshal([]byte(pg_query.Parse("SELECT 1")), &new_parsetree_list)
  if err != nil {
    fmt.Printf("%v\n", err)
  }
  fmt.Printf("%V\n", new_parsetree_list)
  fmt.Printf("%V\n", parsetree_list)

  fmt.Printf("TESTS\n")
  for _, test := range queryTests {
    actual := pg_query.Parse(test.input)
    if actual != test.expected {
      t.Errorf("Parse(%s): expected %s, actual %s", test.input, test.expected, actual)
    }
  }
}
