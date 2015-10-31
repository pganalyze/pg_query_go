// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

func (node CollateClause) MarshalJSON() ([]byte, error) {
	type CollateClauseMarshalAlias CollateClause
	return json.Marshal(map[string]interface{}{
		"COLLATECLAUSE": (*CollateClauseMarshalAlias)(&node),
	})
}

func (node *CollateClause) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}

func (node CollateClause) Deparse() string {
	panic("Not Implemented")
}
