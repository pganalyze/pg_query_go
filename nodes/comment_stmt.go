// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

type CommentStmt struct {
	Objtype ObjectType `json:"objtype"` /* Object's type */
	Objname []Node     `json:"objname"` /* Qualified name of the object */
	Objargs []Node     `json:"objargs"` /* Arguments if needed (eg, for functions) */
	Comment *string    `json:"comment"` /* Comment to insert, or NULL to remove */
}

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
