// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

type ParamListInfoData struct {
	ParamFetchArg  interface{} `json:"paramFetchArg"`
	ParserSetupArg interface{} `json:"parserSetupArg"`
	NumParams      int         `json:"numParams"` /* number of ParamExternDatas following */

}

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
