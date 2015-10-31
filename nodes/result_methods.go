// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

func (node Result) MarshalJSON() ([]byte, error) {
	type ResultMarshalAlias Result
	return json.Marshal(map[string]interface{}{
		"RESULT": (*ResultMarshalAlias)(&node),
	})
}

func (node *Result) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}

func (node Result) Deparse() string {
	panic("Not Implemented")
}
