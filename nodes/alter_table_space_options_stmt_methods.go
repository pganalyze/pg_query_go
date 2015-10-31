// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

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

func (node AlterTableSpaceOptionsStmt) Deparse() string {
	panic("Not Implemented")
}
