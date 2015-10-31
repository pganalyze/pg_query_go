// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

func (node CaseWhen) MarshalJSON() ([]byte, error) {
	type CaseWhenMarshalAlias CaseWhen
	return json.Marshal(map[string]interface{}{
		"WHEN": (*CaseWhenMarshalAlias)(&node),
	})
}

func (node *CaseWhen) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}

func (node CaseWhen) Deparse() string {
	panic("Not Implemented")
}
