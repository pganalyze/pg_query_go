package pg_query

import (
  "encoding/json"
  nodes "github.com/lfittl/pg_query.go/nodes"
)

type ParsetreeList struct {
  Statements []nodes.Node
}

func (input ParsetreeList) MarshalJSON() ([]byte, error) {
  type ParsetreeListAlias ParsetreeList
  return json.Marshal(input.Statements);
}

func (output *ParsetreeList) UnmarshalJSON(input []byte) error {
  var list []json.RawMessage
  json.Unmarshal([]byte(input), &list)

  for _, node := range list {
    output.Statements = append(output.Statements, nodes.UnmarshalNodeJSON(node));
  }

  //fmt.Printf("%v\n", raw_tree)
  // We'd like to process []byte into an array
  /*output.Statements = []nodes.Node{
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
  }*/
  return nil
}
