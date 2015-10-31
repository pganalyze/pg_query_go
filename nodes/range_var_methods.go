// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

func (node RangeVar) MarshalJSON() ([]byte, error) {
	type RangeVarMarshalAlias RangeVar
	return json.Marshal(map[string]interface{}{
		"RANGEVAR": (*RangeVarMarshalAlias)(&node),
	})
}

func (node *RangeVar) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}

func (node RangeVar) Deparse() string {
	panic("Not Implemented")
}
