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

func (output *ParsetreeList) UnmarshalJSON(input []byte) (err error) {
  var list []json.RawMessage
  err = json.Unmarshal([]byte(input), &list)

  for _, node := range list {
    output.Statements = append(output.Statements, nodes.UnmarshalNodeJSON(node));
  }

  return
}
