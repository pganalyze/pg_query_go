// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

func (node Sort) MarshalJSON() ([]byte, error) {
	type SortMarshalAlias Sort
	return json.Marshal(map[string]interface{}{
		"SORT": (*SortMarshalAlias)(&node),
	})
}

func (node *Sort) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}

func (node Sort) Deparse() string {
	panic("Not Implemented")
}
