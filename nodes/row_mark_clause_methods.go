// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

func (node RowMarkClause) MarshalJSON() ([]byte, error) {
	type RowMarkClauseMarshalAlias RowMarkClause
	return json.Marshal(map[string]interface{}{
		"ROWMARKCLAUSE": (*RowMarkClauseMarshalAlias)(&node),
	})
}

func (node *RowMarkClause) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}

func (node RowMarkClause) Deparse() string {
	panic("Not Implemented")
}
