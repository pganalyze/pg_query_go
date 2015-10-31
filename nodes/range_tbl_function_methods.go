// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

func (node RangeTblFunction) MarshalJSON() ([]byte, error) {
	type RangeTblFunctionMarshalAlias RangeTblFunction
	return json.Marshal(map[string]interface{}{
		"RANGETBLFUNCTION": (*RangeTblFunctionMarshalAlias)(&node),
	})
}

func (node *RangeTblFunction) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}

func (node RangeTblFunction) Deparse() string {
	panic("Not Implemented")
}
