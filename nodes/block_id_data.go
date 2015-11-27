// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

type BlockIdData struct {
	BiHi uint16 `json:"bi_hi"`
	BiLo uint16 `json:"bi_lo"`
}

func (node BlockIdData) MarshalJSON() ([]byte, error) {
	type BlockIdDataMarshalAlias BlockIdData
	return json.Marshal(map[string]interface{}{
		"BLOCKIDDATA": (*BlockIdDataMarshalAlias)(&node),
	})
}

func (node *BlockIdData) UnmarshalJSON(input []byte) (err error) {
	var fields map[string]json.RawMessage

	err = json.Unmarshal(input, &fields)
	if err != nil {
		return
	}

	if fields["bi_hi"] != nil {
		err = json.Unmarshal(fields["bi_hi"], &node.BiHi)
		if err != nil {
			return
		}
	}

	if fields["bi_lo"] != nil {
		err = json.Unmarshal(fields["bi_lo"], &node.BiLo)
		if err != nil {
			return
		}
	}

	return
}
