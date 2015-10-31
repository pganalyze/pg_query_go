// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

func (node DefineStmt) MarshalJSON() ([]byte, error) {
	type DefineStmtMarshalAlias DefineStmt
	return json.Marshal(map[string]interface{}{
		"DEFINESTMT": (*DefineStmtMarshalAlias)(&node),
	})
}

func (node *DefineStmt) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}

func (node DefineStmt) Deparse() string {
	panic("Not Implemented")
}
