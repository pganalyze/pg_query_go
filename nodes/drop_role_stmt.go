// Auto-generated from postgres/src/include/nodes/parsenodes.h - DO NOT EDIT

package pg_query

import "encoding/json"

/* ----------------------
 *	Create/Alter/Drop Role Statements
 *
 * Note: these node types are also used for the backwards-compatible
 * Create/Alter/Drop User/Group statements.  In the ALTER and DROP cases
 * there's really no need to distinguish what the original spelling was,
 * but for CREATE we mark the type because the defaults vary.
 * ----------------------
 */
type DropRoleStmt struct {
	Roles     []Node `json:"roles"`      /* List of roles to remove */
	MissingOk bool   `json:"missing_ok"` /* skip error if a role is missing? */
}

func (node DropRoleStmt) MarshalJSON() ([]byte, error) {
	type DropRoleStmtMarshalAlias DropRoleStmt
	return json.Marshal(map[string]interface{}{
		"DROPROLESTMT": (*DropRoleStmtMarshalAlias)(&node),
	})
}

func (node *DropRoleStmt) UnmarshalJSON(input []byte) (err error) {
	var fields map[string]json.RawMessage

	err = json.Unmarshal(input, &fields)
	if err != nil {
		return
	}

	if fields["roles"] != nil {
		node.Roles, err = UnmarshalNodeArrayJSON(fields["roles"])
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
