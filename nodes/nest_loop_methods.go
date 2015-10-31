// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

func (node NestLoop) MarshalJSON() ([]byte, error) {
	type NestLoopMarshalAlias NestLoop
	return json.Marshal(map[string]interface{}{
		"NESTLOOP": (*NestLoopMarshalAlias)(&node),
	})
}

func (node *NestLoop) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}

func (node NestLoop) Deparse() string {
	panic("Not Implemented")
}
