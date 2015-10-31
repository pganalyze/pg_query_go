// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

func (node CreateTrigStmt) MarshalJSON() ([]byte, error) {
	type CreateTrigStmtMarshalAlias CreateTrigStmt
	return json.Marshal(map[string]interface{}{
		"CREATETRIGSTMT": (*CreateTrigStmtMarshalAlias)(&node),
	})
}

func (node *CreateTrigStmt) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}

func (node CreateTrigStmt) Deparse() string {
	panic("Not Implemented")
}
