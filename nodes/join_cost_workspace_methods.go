// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

func (node JoinCostWorkspace) MarshalJSON() ([]byte, error) {
	type JoinCostWorkspaceMarshalAlias JoinCostWorkspace
	return json.Marshal(map[string]interface{}{
		"JOINCOSTWORKSPACE": (*JoinCostWorkspaceMarshalAlias)(&node),
	})
}

func (node *JoinCostWorkspace) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}

func (node JoinCostWorkspace) Deparse() string {
	panic("Not Implemented")
}
