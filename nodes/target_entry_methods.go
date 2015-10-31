// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

func (node TargetEntry) MarshalJSON() ([]byte, error) {
	type TargetEntryMarshalAlias TargetEntry
	return json.Marshal(map[string]interface{}{
		"TARGETENTRY": (*TargetEntryMarshalAlias)(&node),
	})
}

func (node *TargetEntry) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}

func (node TargetEntry) Deparse() string {
	panic("Not Implemented")
}
