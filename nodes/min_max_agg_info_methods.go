// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

func (node MinMaxAggInfo) MarshalJSON() ([]byte, error) {
	type MinMaxAggInfoMarshalAlias MinMaxAggInfo
	return json.Marshal(map[string]interface{}{
		"MINMAXAGGINFO": (*MinMaxAggInfoMarshalAlias)(&node),
	})
}

func (node *MinMaxAggInfo) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}

func (node MinMaxAggInfo) Deparse() string {
	panic("Not Implemented")
}
