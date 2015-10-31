// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

func (node ColumnDef) MarshalJSON() ([]byte, error) {
	type ColumnDefMarshalAlias ColumnDef
	return json.Marshal(map[string]interface{}{
		"COLUMNDEF": (*ColumnDefMarshalAlias)(&node),
	})
}

func (node *ColumnDef) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}

func (node ColumnDef) Deparse() string {
	panic("Not Implemented")
}
