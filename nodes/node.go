package pg_query

import "encoding/json"
import "fmt"
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

func UnmarshalNodeFieldJSON(input []byte, node Node) (err error) {
	var fields map[string]json.RawMessage

	err = json.Unmarshal([]byte(input), &fields)
	if err != nil {
		return
	}

	s := reflect.ValueOf(node).Elem()
	typeOfT := s.Type()

	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		f2 := typeOfT.Field(i)
		fieldName := typeOfT.Field(i).Tag.Get("json")
		nodeType := reflect.TypeOf((*Node)(nil)).Elem()

		// Null is the default, so we just skip the type magic here
		if string(fields[fieldName]) == "null" {
			continue
		}

		switch f.Interface().(type) {
		case int:
			var value int64
			json.Unmarshal(fields[fieldName], &value)
			f.SetInt(value)
		case Value:
			// FIXME: We shouldn't special case like this
			var value Value
			json.Unmarshal(fields[fieldName], &value)
			f.Set(reflect.ValueOf(value))
		case *string:
			var value string
			json.Unmarshal(fields[fieldName], &value)
			f.Set(reflect.ValueOf(&value))
		case []Node:
			var list []json.RawMessage
			var nodes []Node
			json.Unmarshal(fields[fieldName], &list)
			for _, node := range list {
				nodes = append(nodes, UnmarshalNodeJSON(node))
			}
			f.Set(reflect.ValueOf(nodes))
		default:
			if f2.Type.Implements(nodeType) {
				node := UnmarshalNodeJSON(fields[fieldName])
				if node != nil {
					f.Set(reflect.ValueOf(node))
				}
			} else {
				err = fmt.Errorf("Invalid type %v", f.Type())
			}
		}
	}

	return
}
