// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

func (node RelOptInfo) MarshalJSON() ([]byte, error) {
	type RelOptInfoMarshalAlias RelOptInfo
	return json.Marshal(map[string]interface{}{
		"RELOPTINFO": (*RelOptInfoMarshalAlias)(&node),
	})
}

func (node *RelOptInfo) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}

func (node RelOptInfo) Deparse() string {
	panic("Not Implemented")
}
