// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

type AlterTableMoveAllStmt struct {
	OrigTablespacename *string    `json:"orig_tablespacename"`
	Objtype            ObjectType `json:"objtype"` /* Object type to move */
	Roles              []Node     `json:"roles"`   /* List of roles to move objects of */
	NewTablespacename  *string    `json:"new_tablespacename"`
	Nowait             bool       `json:"nowait"`
}

func (node AlterTableMoveAllStmt) MarshalJSON() ([]byte, error) {
	type AlterTableMoveAllStmtMarshalAlias AlterTableMoveAllStmt
	return json.Marshal(map[string]interface{}{
		"ALTERTABLEMOVEALLSTMT": (*AlterTableMoveAllStmtMarshalAlias)(&node),
	})
}

func (node *AlterTableMoveAllStmt) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}
