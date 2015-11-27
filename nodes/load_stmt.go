// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

type LoadStmt struct {
	Filename *string `json:"filename"` /* file to load */
}

func (node LoadStmt) MarshalJSON() ([]byte, error) {
	type LoadStmtMarshalAlias LoadStmt
	return json.Marshal(map[string]interface{}{
		"LOADSTMT": (*LoadStmtMarshalAlias)(&node),
	})
}

func (node *LoadStmt) UnmarshalJSON(input []byte) (err error) {
	var fields map[string]json.RawMessage

	err = json.Unmarshal(input, &fields)
	if err != nil {
		return
	}

	if fields["filename"] != nil {
		err = json.Unmarshal(fields["filename"], &node.Filename)
		if err != nil {
			return
		}
	}

	return
}
