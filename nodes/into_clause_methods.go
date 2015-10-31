// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

func (node IntoClause) MarshalJSON() ([]byte, error) {
	type IntoClauseMarshalAlias IntoClause
	return json.Marshal(map[string]interface{}{
		"INTOCLAUSE": (*IntoClauseMarshalAlias)(&node),
	})
}

func (node *IntoClause) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}

func (node IntoClause) Deparse() string {
	panic("Not Implemented")
}
