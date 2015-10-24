// Auto-generated - DO NOT EDIT

package pg_query

type AlterTableMoveAllStmt struct {
	OrigTablespacename *string    `json:"orig_tablespacename"`
	Objtype            ObjectType `json:"objtype"` /* Object type to move */
	Roles              []Node     `json:"roles"`   /* List of roles to move objects of */
	NewTablespacename  *string    `json:"new_tablespacename"`
	Nowait             bool       `json:"nowait"`
}
