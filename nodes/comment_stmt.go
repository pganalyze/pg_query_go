// Auto-generated - DO NOT EDIT

package pg_query

type CommentStmt struct {
	Objtype ObjectType `json:"objtype"` /* Object's type */
	Objname []Node     `json:"objname"` /* Qualified name of the object */
	Objargs []Node     `json:"objargs"` /* Arguments if needed (eg, for functions) */
	Comment *string    `json:"comment"` /* Comment to insert, or NULL to remove */
}
