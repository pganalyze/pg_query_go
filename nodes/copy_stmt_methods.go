// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

func (node CopyStmt) MarshalJSON() ([]byte, error) {
	type CopyStmtMarshalAlias CopyStmt
	return json.Marshal(map[string]interface{}{
		"COPY": (*CopyStmtMarshalAlias)(&node),
	})
}

func (node *CopyStmt) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}

func (node CopyStmt) Deparse() string {
	panic("Not Implemented")
}
