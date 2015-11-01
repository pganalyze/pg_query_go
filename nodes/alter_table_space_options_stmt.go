// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

type AlterTableSpaceOptionsStmt struct {
	Tablespacename *string `json:"tablespacename"`
	Options        []Node  `json:"options"`
	IsReset        bool    `json:"isReset"`
}

func (node AlterTableSpaceOptionsStmt) MarshalJSON() ([]byte, error) {
	type AlterTableSpaceOptionsStmtMarshalAlias AlterTableSpaceOptionsStmt
	return json.Marshal(map[string]interface{}{
		"ALTERTABLESPACEOPTIONSSTMT": (*AlterTableSpaceOptionsStmtMarshalAlias)(&node),
	})
}

func (node *AlterTableSpaceOptionsStmt) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}
