// Auto-generated from postgres/src/include/nodes/params.h - DO NOT EDIT

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
	var fields map[string]json.RawMessage

	err = json.Unmarshal(input, &fields)
	if err != nil {
		return
	}

	if fields["paramFetchArg"] != nil {
		err = json.Unmarshal(fields["paramFetchArg"], &node.ParamFetchArg)
		if err != nil {
			return
		}
	}

	if fields["parserSetupArg"] != nil {
		err = json.Unmarshal(fields["parserSetupArg"], &node.ParserSetupArg)
		if err != nil {
			return
		}
	}

	if fields["numParams"] != nil {
		err = json.Unmarshal(fields["numParams"], &node.NumParams)
		if err != nil {
			return
		}
	}

	return
}
