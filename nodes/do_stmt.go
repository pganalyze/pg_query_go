// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

type DoStmt struct {
	Args []Node `json:"args"` /* List of DefElem nodes */
}

func (node DoStmt) MarshalJSON() ([]byte, error) {
	type DoStmtMarshalAlias DoStmt
	return json.Marshal(map[string]interface{}{
		"DOSTMT": (*DoStmtMarshalAlias)(&node),
	})
}

func (node *DoStmt) UnmarshalJSON(input []byte) (err error) {
	var fields map[string]json.RawMessage

	err = json.Unmarshal(input, &fields)
	if err != nil {
		return
	}

	if fields["args"] != nil {
		node.Args, err = UnmarshalNodeArrayJSON(fields["args"])
		if err != nil {
			return
		}
	}

	return
}
