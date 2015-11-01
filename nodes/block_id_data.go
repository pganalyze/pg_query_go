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
	err = UnmarshalNodeFieldJSON(input, node)
	return
}
