// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

func (node PrivGrantee) MarshalJSON() ([]byte, error) {
	type PrivGranteeMarshalAlias PrivGrantee
	return json.Marshal(map[string]interface{}{
		"PRIVGRANTEE": (*PrivGranteeMarshalAlias)(&node),
	})
}

func (node *PrivGrantee) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}

func (node PrivGrantee) Deparse() string {
	panic("Not Implemented")
}
