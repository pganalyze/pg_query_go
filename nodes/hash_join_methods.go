// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

func (node HashJoin) MarshalJSON() ([]byte, error) {
	type HashJoinMarshalAlias HashJoin
	return json.Marshal(map[string]interface{}{
		"HASHJOIN": (*HashJoinMarshalAlias)(&node),
	})
}

func (node *HashJoin) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}

func (node HashJoin) Deparse() string {
	panic("Not Implemented")
}
