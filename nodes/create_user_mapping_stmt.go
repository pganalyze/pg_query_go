// Auto-generated from postgres/src/include/nodes/parsenodes.h - DO NOT EDIT

package pg_query

import "encoding/json"

/* ----------------------
 *		Create/Drop USER MAPPING Statements
 * ----------------------
 */
type CreateUserMappingStmt struct {
	Username   *string `json:"username"`   /* username or PUBLIC/CURRENT_USER */
	Servername *string `json:"servername"` /* server name */
	Options    []Node  `json:"options"`    /* generic options to server */
}

func (node CreateUserMappingStmt) MarshalJSON() ([]byte, error) {
	type CreateUserMappingStmtMarshalAlias CreateUserMappingStmt
	return json.Marshal(map[string]interface{}{
		"CreateUserMappingStmt": (*CreateUserMappingStmtMarshalAlias)(&node),
	})
}

func (node *CreateUserMappingStmt) UnmarshalJSON(input []byte) (err error) {
	var fields map[string]json.RawMessage

	err = json.Unmarshal(input, &fields)
	if err != nil {
		return
	}

	if fields["username"] != nil {
		err = json.Unmarshal(fields["username"], &node.Username)
		if err != nil {
			return
		}
	}

	if fields["servername"] != nil {
		err = json.Unmarshal(fields["servername"], &node.Servername)
		if err != nil {
			return
		}
	}

	if fields["options"] != nil {
		node.Options, err = UnmarshalNodeArrayJSON(fields["options"])
		if err != nil {
			return
		}
	}

	return
}
