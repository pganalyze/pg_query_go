// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

type NotifyStmt struct {
	Conditionname *string `json:"conditionname"` /* condition name to notify */
	Payload       *string `json:"payload"`       /* the payload string, or NULL if none */
}

func (node NotifyStmt) MarshalJSON() ([]byte, error) {
	type NotifyStmtMarshalAlias NotifyStmt
	return json.Marshal(map[string]interface{}{
		"NOTIFYSTMT": (*NotifyStmtMarshalAlias)(&node),
	})
}

func (node *NotifyStmt) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}
