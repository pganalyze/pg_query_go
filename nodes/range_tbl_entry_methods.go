// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

func (node RangeTblEntry) MarshalJSON() ([]byte, error) {
	type RangeTblEntryMarshalAlias RangeTblEntry
	return json.Marshal(map[string]interface{}{
		"RANGETBLENTRY": (*RangeTblEntryMarshalAlias)(&node),
	})
}

func (node *RangeTblEntry) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}

func (node RangeTblEntry) Deparse() string {
	panic("Not Implemented")
}
