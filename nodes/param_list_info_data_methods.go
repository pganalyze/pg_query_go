// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

func (node ParamListInfoData) MarshalJSON() ([]byte, error) {
	type ParamListInfoDataMarshalAlias ParamListInfoData
	return json.Marshal(map[string]interface{}{
		"PARAMLISTINFODATA": (*ParamListInfoDataMarshalAlias)(&node),
	})
}

func (node *ParamListInfoData) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}

func (node ParamListInfoData) Deparse() string {
	panic("Not Implemented")
}
