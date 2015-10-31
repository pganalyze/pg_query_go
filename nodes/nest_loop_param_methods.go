// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

func (node NestLoopParam) MarshalJSON() ([]byte, error) {
	type NestLoopParamMarshalAlias NestLoopParam
	return json.Marshal(map[string]interface{}{
		"NESTLOOPPARAM": (*NestLoopParamMarshalAlias)(&node),
	})
}

func (node *NestLoopParam) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}

func (node NestLoopParam) Deparse() string {
	panic("Not Implemented")
}
