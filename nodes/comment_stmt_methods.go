// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

func (node CommentStmt) MarshalJSON() ([]byte, error) {
	type CommentStmtMarshalAlias CommentStmt
	return json.Marshal(map[string]interface{}{
		"COMMENTSTMT": (*CommentStmtMarshalAlias)(&node),
	})
}

func (node *CommentStmt) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}

func (node CommentStmt) Deparse() string {
	panic("Not Implemented")
}
