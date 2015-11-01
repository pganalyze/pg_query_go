// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

type NestLoopParam struct {
	Paramno  int  `json:"paramno"`  /* number of the PARAM_EXEC Param to set */
	Paramval *Var `json:"paramval"` /* outer-relation Var to assign to Param */
}

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
