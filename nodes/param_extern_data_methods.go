// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

func (node ParamExternData) MarshalJSON() ([]byte, error) {
	type ParamExternDataMarshalAlias ParamExternData
	return json.Marshal(map[string]interface{}{
		"PARAMEXTERNDATA": (*ParamExternDataMarshalAlias)(&node),
	})
}

func (node *ParamExternData) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}

func (node ParamExternData) Deparse() string {
	panic("Not Implemented")
}
