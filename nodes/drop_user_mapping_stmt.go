// Auto-generated from postgres/src/include/nodes/parsenodes.h - DO NOT EDIT

package pg_query

import "encoding/json"

/* ----------------------
 *		Create/Drop USER MAPPING Statements
 * ----------------------
 */
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

	if fields["missing_ok"] != nil {
		err = json.Unmarshal(fields["missing_ok"], &node.MissingOk)
		if err != nil {
			return
		}
	}

	return
}
