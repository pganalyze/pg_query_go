// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

func (node MergeJoin) MarshalJSON() ([]byte, error) {
	type MergeJoinMarshalAlias MergeJoin
	return json.Marshal(map[string]interface{}{
		"MERGEJOIN": (*MergeJoinMarshalAlias)(&node),
	})
}

func (node *MergeJoin) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}

func (node MergeJoin) Deparse() string {
	panic("Not Implemented")
}
