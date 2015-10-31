// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

func (node DropUserMappingStmt) MarshalJSON() ([]byte, error) {
	type DropUserMappingStmtMarshalAlias DropUserMappingStmt
	return json.Marshal(map[string]interface{}{
		"DROPUSERMAPPINGSTMT": (*DropUserMappingStmtMarshalAlias)(&node),
	})
}

func (node *DropUserMappingStmt) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}

func (node DropUserMappingStmt) Deparse() string {
	panic("Not Implemented")
}
