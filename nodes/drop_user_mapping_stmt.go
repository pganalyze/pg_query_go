// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

type DropUserMappingStmt struct {
	Username   *string `json:"username"`   /* username or PUBLIC/CURRENT_USER */
	Servername *string `json:"servername"` /* server name */
	MissingOk  bool    `json:"missing_ok"` /* ignore missing mappings */
}

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
