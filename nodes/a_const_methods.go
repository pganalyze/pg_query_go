package pg_query

import "encoding/json"

func (input A_Const) MarshalJSON() ([]byte, error) {
	type A_ConstAlias A_Const
	return json.Marshal(map[string]interface{}{
		"A_CONST": (*A_ConstAlias)(&input),
	})
}

func (aConst *A_Const) UnmarshalJSON(input []byte) error {
	return UnmarshalNodeFieldJSON(input, aConst)
}

func (aConst A_Const) Deparse() string {
	panic("Not Implemented")
}
