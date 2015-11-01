// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

type ExplainStmt struct {
	Query   Node   `json:"query"`   /* the query (see comments above) */
	Options []Node `json:"options"` /* list of DefElem nodes */
}

func (node ExplainStmt) MarshalJSON() ([]byte, error) {
	type ExplainStmtMarshalAlias ExplainStmt
	return json.Marshal(map[string]interface{}{
		"EXPLAIN": (*ExplainStmtMarshalAlias)(&node),
	})
}

func (node *ExplainStmt) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}
