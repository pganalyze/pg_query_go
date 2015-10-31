// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

func (node Group) MarshalJSON() ([]byte, error) {
	type GroupMarshalAlias Group
	return json.Marshal(map[string]interface{}{
		"GROUP": (*GroupMarshalAlias)(&node),
	})
}

func (node *Group) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}

func (node Group) Deparse() string {
	panic("Not Implemented")
}
