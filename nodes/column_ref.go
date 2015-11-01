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
	err = UnmarshalNodeFieldJSON(input, node)
	return
}
