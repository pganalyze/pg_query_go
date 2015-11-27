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
	var fields map[string]json.RawMessage

	err = json.Unmarshal(input, &fields)
	if err != nil {
		return
	}

	if fields["value"] != nil {
		err = json.Unmarshal(fields["value"], &node.Value)
		if err != nil {
			return
		}
	}

	if fields["isnull"] != nil {
		err = json.Unmarshal(fields["isnull"], &node.Isnull)
		if err != nil {
			return
		}
	}

	if fields["pflags"] != nil {
		err = json.Unmarshal(fields["pflags"], &node.Pflags)
		if err != nil {
			return
		}
	}

	if fields["ptype"] != nil {
		err = json.Unmarshal(fields["ptype"], &node.Ptype)
		if err != nil {
			return
		}
	}

	return
}
