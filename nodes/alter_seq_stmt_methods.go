// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

func (node AlterSeqStmt) MarshalJSON() ([]byte, error) {
	type AlterSeqStmtMarshalAlias AlterSeqStmt
	return json.Marshal(map[string]interface{}{
		"ALTERSEQSTMT": (*AlterSeqStmtMarshalAlias)(&node),
	})
}

func (node *AlterSeqStmt) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}

func (node AlterSeqStmt) Deparse() string {
	panic("Not Implemented")
}
