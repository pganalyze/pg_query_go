// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

func (node EquivalenceMember) MarshalJSON() ([]byte, error) {
	type EquivalenceMemberMarshalAlias EquivalenceMember
	return json.Marshal(map[string]interface{}{
		"EQUIVALENCEMEMBER": (*EquivalenceMemberMarshalAlias)(&node),
	})
}

func (node *EquivalenceMember) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}

func (node EquivalenceMember) Deparse() string {
	panic("Not Implemented")
}
