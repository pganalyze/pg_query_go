// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

func (node SortBy) MarshalJSON() ([]byte, error) {
	type SortByMarshalAlias SortBy
	return json.Marshal(map[string]interface{}{
		"SORTBY": (*SortByMarshalAlias)(&node),
	})
}

func (node *SortBy) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}

func (node SortBy) Deparse() string {
	panic("Not Implemented")
}
