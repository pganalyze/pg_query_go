// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

type CheckPointStmt struct {
}

func (node CheckPointStmt) MarshalJSON() ([]byte, error) {
	type CheckPointStmtMarshalAlias CheckPointStmt
	return json.Marshal(map[string]interface{}{
		"CHECKPOINT": (*CheckPointStmtMarshalAlias)(&node),
	})
}

func (node *CheckPointStmt) UnmarshalJSON(input []byte) (err error) {
	var fields map[string]json.RawMessage

	err = json.Unmarshal(input, &fields)
	if err != nil {
		return
	}

	return
}
