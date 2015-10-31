// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

func (node PrepareStmt) MarshalJSON() ([]byte, error) {
	type PrepareStmtMarshalAlias PrepareStmt
	return json.Marshal(map[string]interface{}{
		"PREPARESTMT": (*PrepareStmtMarshalAlias)(&node),
	})
}

func (node *PrepareStmt) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}

func (node PrepareStmt) Deparse() string {
	panic("Not Implemented")
}
