// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

type RangeTblRef struct {
	Rtindex int `json:"rtindex"`
}

func (node RangeTblRef) MarshalJSON() ([]byte, error) {
	type RangeTblRefMarshalAlias RangeTblRef
	return json.Marshal(map[string]interface{}{
		"RANGETBLREF": (*RangeTblRefMarshalAlias)(&node),
	})
}

func (node *RangeTblRef) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}
