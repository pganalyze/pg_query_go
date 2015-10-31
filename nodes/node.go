package pg_query

import "encoding/json"
import "fmt"
import "reflect"

type Node interface {
	Deparse() string
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
		if fields[fieldName] == nil || string(fields[fieldName]) == "null" {
			continue
		}

		switch f.Interface().(type) {
		case int:
			var value int64
			err = json.Unmarshal(fields[fieldName], &value)
			if err != nil {
				return
			}
			f.SetInt(value)
		case byte:
			var value string
			err = json.Unmarshal(fields[fieldName], &value)
			if err != nil {
				return
			}
			f.Set(reflect.ValueOf([]byte(value)[0]))
		case *string:
			var value string
			err = json.Unmarshal(fields[fieldName], &value)
			if err != nil {
				return
			}
			f.Set(reflect.ValueOf(&value))
		case []Node:
			var list []json.RawMessage
			var nodes []Node
			err = json.Unmarshal(fields[fieldName], &list)
			if err != nil {
				return
			}
			for _, nodeJson := range list {
				node, err = UnmarshalNodeJSON(nodeJson)
				if err != nil {
					return
				}
				nodes = append(nodes, node)
			}
			f.Set(reflect.ValueOf(nodes))
		default:
			if f2.Type.Kind() == reflect.Uint {
				var enumValue uint
				err = json.Unmarshal(fields[fieldName], &enumValue)
				if err != nil {
					return
				}
				f.SetUint(uint64(enumValue))
				continue
			}

			if !f2.Type.Implements(nodeType) {
				err = fmt.Errorf("Invalid type %v", f.Type())
				return
			}

			node, err = UnmarshalNodeJSON(fields[fieldName])

			if err != nil {
				return
			}

			if node != nil {
				f.Set(reflect.ValueOf(node))
			}
		}
	}

	return
}
