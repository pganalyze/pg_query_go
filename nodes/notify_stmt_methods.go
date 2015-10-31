// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

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

func (node NotifyStmt) Deparse() string {
	panic("Not Implemented")
}
