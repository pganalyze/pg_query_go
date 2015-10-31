// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

func (node SecLabelStmt) MarshalJSON() ([]byte, error) {
	type SecLabelStmtMarshalAlias SecLabelStmt
	return json.Marshal(map[string]interface{}{
		"SECLABELSTMT": (*SecLabelStmtMarshalAlias)(&node),
	})
}

func (node *SecLabelStmt) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}

func (node SecLabelStmt) Deparse() string {
	panic("Not Implemented")
}
