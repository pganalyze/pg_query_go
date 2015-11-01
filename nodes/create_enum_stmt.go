// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

type CreateEnumStmt struct {
	TypeName []Node `json:"typeName"` /* qualified name (list of Value strings) */
	Vals     []Node `json:"vals"`     /* enum values (list of Value strings) */
}

func (node CreateEnumStmt) MarshalJSON() ([]byte, error) {
	type CreateEnumStmtMarshalAlias CreateEnumStmt
	return json.Marshal(map[string]interface{}{
		"CREATEENUMSTMT": (*CreateEnumStmtMarshalAlias)(&node),
	})
}

func (node *CreateEnumStmt) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}
