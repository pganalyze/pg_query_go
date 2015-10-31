// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

func (node ParamExecData) MarshalJSON() ([]byte, error) {
	type ParamExecDataMarshalAlias ParamExecData
	return json.Marshal(map[string]interface{}{
		"PARAMEXECDATA": (*ParamExecDataMarshalAlias)(&node),
	})
}

func (node *ParamExecData) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}

func (node ParamExecData) Deparse() string {
	panic("Not Implemented")
}
