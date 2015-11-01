// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

type ParamExternData struct {
	Value  Datum  `json:"value"`  /* parameter value */
	Isnull bool   `json:"isnull"` /* is it NULL? */
	Pflags uint16 `json:"pflags"` /* flag bits, see above */
	Ptype  Oid    `json:"ptype"`  /* parameter's datatype, or 0 */
}

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
