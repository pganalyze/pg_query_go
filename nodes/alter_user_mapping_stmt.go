// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

type AlterUserMappingStmt struct {
	Username   *string `json:"username"`   /* username or PUBLIC/CURRENT_USER */
	Servername *string `json:"servername"` /* server name */
	Options    []Node  `json:"options"`    /* generic options to server */
}

func (node AlterUserMappingStmt) MarshalJSON() ([]byte, error) {
	type AlterUserMappingStmtMarshalAlias AlterUserMappingStmt
	return json.Marshal(map[string]interface{}{
		"ALTERUSERMAPPINGSTMT": (*AlterUserMappingStmtMarshalAlias)(&node),
	})
}

func (node *AlterUserMappingStmt) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}
