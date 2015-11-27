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
	var fields map[string]json.RawMessage

	err = json.Unmarshal(input, &fields)
	if err != nil {
		return
	}

	if fields["query"] != nil {
		node.Query, err = UnmarshalNodeJSON(fields["query"])
		if err != nil {
			return
		}
	}

	if fields["options"] != nil {
		node.Options, err = UnmarshalNodeArrayJSON(fields["options"])
		if err != nil {
			return
		}
	}

	return
}
