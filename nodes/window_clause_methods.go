// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

func (node WindowClause) MarshalJSON() ([]byte, error) {
	type WindowClauseMarshalAlias WindowClause
	return json.Marshal(map[string]interface{}{
		"WINDOWCLAUSE": (*WindowClauseMarshalAlias)(&node),
	})
}

func (node *WindowClause) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}

func (node WindowClause) Deparse() string {
	panic("Not Implemented")
}
