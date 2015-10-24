// Auto-generated - DO NOT EDIT

package pg_query

type AlterExtensionContentsStmt struct {
	Extname *string    `json:"extname"` /* Extension's name */
	Action  int        `json:"action"`  /* +1 = add object, -1 = drop object */
	Objtype ObjectType `json:"objtype"` /* Object's type */
	Objname []Node     `json:"objname"` /* Qualified name of the object */
	Objargs []Node     `json:"objargs"` /* Arguments if needed (eg, for functions) */
}
