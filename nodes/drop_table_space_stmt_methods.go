// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

func (node DropTableSpaceStmt) MarshalJSON() ([]byte, error) {
	type DropTableSpaceStmtMarshalAlias DropTableSpaceStmt
	return json.Marshal(map[string]interface{}{
		"DROPTABLESPACESTMT": (*DropTableSpaceStmtMarshalAlias)(&node),
	})
}

func (node *DropTableSpaceStmt) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}

func (node DropTableSpaceStmt) Deparse() string {
	panic("Not Implemented")
}
