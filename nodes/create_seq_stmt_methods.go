// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

func (node CreateSeqStmt) MarshalJSON() ([]byte, error) {
	type CreateSeqStmtMarshalAlias CreateSeqStmt
	return json.Marshal(map[string]interface{}{
		"CREATESEQSTMT": (*CreateSeqStmtMarshalAlias)(&node),
	})
}

func (node *CreateSeqStmt) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}

func (node CreateSeqStmt) Deparse() string {
	panic("Not Implemented")
}
