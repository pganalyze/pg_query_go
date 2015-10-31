// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

func (node varatt_external) MarshalJSON() ([]byte, error) {
	type varatt_externalMarshalAlias varatt_external
	return json.Marshal(map[string]interface{}{
		"VARATT_EXTERNAL": (*varatt_externalMarshalAlias)(&node),
	})
}

func (node *varatt_external) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}

func (node varatt_external) Deparse() string {
	panic("Not Implemented")
}
