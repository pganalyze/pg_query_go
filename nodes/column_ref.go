// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

type ColumnRef struct {
	Fields   []Node `json:"fields"`   /* field names (Value strings) or A_Star */
	Location int    `json:"location"` /* token location, or -1 if unknown */
}

func (node ColumnRef) MarshalJSON() ([]byte, error) {
	type ColumnRefMarshalAlias ColumnRef
	return json.Marshal(map[string]interface{}{
		"COLUMNREF": (*ColumnRefMarshalAlias)(&node),
	})
}

func (node *ColumnRef) UnmarshalJSON(input []byte) (err error) {
	var fields map[string]json.RawMessage

	err = json.Unmarshal(input, &fields)
	if err != nil {
		return
	}

	if fields["fields"] != nil {
		node.Fields, err = UnmarshalNodeArrayJSON(fields["fields"])
		if err != nil {
			return
		}
	}

	if fields["location"] != nil {
		err = json.Unmarshal(fields["location"], &node.Location)
		if err != nil {
			return
		}
	}

	return
}
