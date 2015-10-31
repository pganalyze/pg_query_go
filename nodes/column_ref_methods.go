// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

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

func (node ColumnRef) Deparse() string {
	panic("Not Implemented")
}
