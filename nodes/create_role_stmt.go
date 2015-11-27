// Auto-generated from postgres/src/include/nodes/parsenodes.h - DO NOT EDIT

package pg_query

import "encoding/json"

type CreateRoleStmt struct {
	StmtType RoleStmtType `json:"stmt_type"` /* ROLE/USER/GROUP */
	Role     *string      `json:"role"`      /* role name */
	Options  []Node       `json:"options"`   /* List of DefElem nodes */
}

func (node CreateRoleStmt) MarshalJSON() ([]byte, error) {
	type CreateRoleStmtMarshalAlias CreateRoleStmt
	return json.Marshal(map[string]interface{}{
		"CREATEROLESTMT": (*CreateRoleStmtMarshalAlias)(&node),
	})
}

func (node *CreateRoleStmt) UnmarshalJSON(input []byte) (err error) {
	var fields map[string]json.RawMessage

	err = json.Unmarshal(input, &fields)
	if err != nil {
		return
	}

	if fields["stmt_type"] != nil {
		err = json.Unmarshal(fields["stmt_type"], &node.StmtType)
		if err != nil {
			return
		}
	}

	if fields["role"] != nil {
		err = json.Unmarshal(fields["role"], &node.Role)
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
