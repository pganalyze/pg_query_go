// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

type PrivGrantee struct {
	Rolname *string `json:"rolname"` /* if NULL then PUBLIC */
}

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
