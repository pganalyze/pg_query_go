// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

type CreateUserMappingStmt struct {
	Username   *string `json:"username"`   /* username or PUBLIC/CURRENT_USER */
	Servername *string `json:"servername"` /* server name */
	Options    []Node  `json:"options"`    /* generic options to server */
}

func (node CreateUserMappingStmt) MarshalJSON() ([]byte, error) {
	type CreateUserMappingStmtMarshalAlias CreateUserMappingStmt
	return json.Marshal(map[string]interface{}{
		"CREATEUSERMAPPINGSTMT": (*CreateUserMappingStmtMarshalAlias)(&node),
	})
}

func (node *CreateUserMappingStmt) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}
