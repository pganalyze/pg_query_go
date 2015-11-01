// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

type WithClause struct {
	Ctes      []Node `json:"ctes"`      /* list of CommonTableExprs */
	Recursive bool   `json:"recursive"` /* true = WITH RECURSIVE */
	Location  int    `json:"location"`  /* token location, or -1 if unknown */
}

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
