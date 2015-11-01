// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

type TableLikeClause struct {
	Relation *RangeVar `json:"relation"`
	Options  uint32    `json:"options"` /* OR of TableLikeOption flags */
}

func (node TableLikeClause) MarshalJSON() ([]byte, error) {
	type TableLikeClauseMarshalAlias TableLikeClause
	return json.Marshal(map[string]interface{}{
		"TABLELIKECLAUSE": (*TableLikeClauseMarshalAlias)(&node),
	})
}

func (node *TableLikeClause) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}
