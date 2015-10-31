// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

func (node WithClause) MarshalJSON() ([]byte, error) {
	type WithClauseMarshalAlias WithClause
	return json.Marshal(map[string]interface{}{
		"WITHCLAUSE": (*WithClauseMarshalAlias)(&node),
	})
}

func (node *WithClause) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}

func (node WithClause) Deparse() string {
	panic("Not Implemented")
}
