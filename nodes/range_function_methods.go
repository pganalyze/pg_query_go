// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

func (node RangeFunction) MarshalJSON() ([]byte, error) {
	type RangeFunctionMarshalAlias RangeFunction
	return json.Marshal(map[string]interface{}{
		"RANGEFUNCTION": (*RangeFunctionMarshalAlias)(&node),
	})
}

func (node *RangeFunction) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}

func (node RangeFunction) Deparse() string {
	panic("Not Implemented")
}
