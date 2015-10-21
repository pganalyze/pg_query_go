package pg_query

import "encoding/json"
import "reflect"

type Node interface {
  Deparse() string
}

func UnmarshalNodeJSON(input json.RawMessage) Node {
  var nodeMap map[string]json.RawMessage
  json.Unmarshal(input, &nodeMap)

  for nodeType, jsonText := range nodeMap {
    switch nodeType {
    case "SELECT":
      var selectStmt SelectStmt
      json.Unmarshal(jsonText, &selectStmt)
      return selectStmt
    case "RESTARGET":
      var resTarget ResTarget
      json.Unmarshal(jsonText, &resTarget)
      return resTarget
    case "A_CONST":
      var aConst A_Const
      json.Unmarshal(jsonText, &aConst)
      return aConst
    }
  }

  return nil
}

func UnmarshalNodeFieldJSON(input []byte, node Node) error {
  var fields map[string]json.RawMessage
  json.Unmarshal([]byte(input), &fields)

  s := reflect.ValueOf(node).Elem()
  typeOfT := s.Type()
  for i := 0; i < s.NumField(); i++ {
      f := s.Field(i)
      f2 := typeOfT.Field(i)
      fieldName := typeOfT.Field(i).Tag.Get("json")// || typeOfT.Field(i).Name

      nodeType := reflect.TypeOf((*Node)(nil)).Elem()

      if (f.Type() == reflect.TypeOf([]Node{})) {
        var list []json.RawMessage
        var nodes []Node
        json.Unmarshal(fields[fieldName], &list)
        for _, node := range list {
          nodes = append(nodes, UnmarshalNodeJSON(node))
        }
        f.Set(reflect.ValueOf(nodes))
      } else if (f2.Type.Implements(nodeType)) {
        node := UnmarshalNodeJSON(fields[fieldName])
        if (node != nil) {
          f.Set(reflect.ValueOf(node))
        }
      } else {
        // TODO: Fix this to work correctly
        var dummy int64
        json.Unmarshal(fields[fieldName], &dummy)
        //f.SetInt(dummy)
      }
  }

  return nil
}
