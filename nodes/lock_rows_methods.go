// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

func (node LockRows) MarshalJSON() ([]byte, error) {
	type LockRowsMarshalAlias LockRows
	return json.Marshal(map[string]interface{}{
		"LOCKROWS": (*LockRowsMarshalAlias)(&node),
	})
}

func (node *LockRows) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}

func (node LockRows) Deparse() string {
	panic("Not Implemented")
}
