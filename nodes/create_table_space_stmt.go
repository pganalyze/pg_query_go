// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

type CreateTableSpaceStmt struct {
	Tablespacename *string `json:"tablespacename"`
	Owner          *string `json:"owner"`
	Location       *string `json:"location"`
	Options        []Node  `json:"options"`
}

func (node CreateTableSpaceStmt) MarshalJSON() ([]byte, error) {
	type CreateTableSpaceStmtMarshalAlias CreateTableSpaceStmt
	return json.Marshal(map[string]interface{}{
		"CREATETABLESPACESTMT": (*CreateTableSpaceStmtMarshalAlias)(&node),
	})
}

func (node *CreateTableSpaceStmt) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}
