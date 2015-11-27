// Auto-generated from postgres/src/include/nodes/parsenodes.h - DO NOT EDIT

package pg_query

import "encoding/json"

type AlterRoleStmt struct {
	Role    *string `json:"role"`    /* role name */
	Options []Node  `json:"options"` /* List of DefElem nodes */
	Action  int     `json:"action"`  /* +1 = add members, -1 = drop members */
}

func (node AlterRoleStmt) MarshalJSON() ([]byte, error) {
	type AlterRoleStmtMarshalAlias AlterRoleStmt
	return json.Marshal(map[string]interface{}{
		"ALTERROLESTMT": (*AlterRoleStmtMarshalAlias)(&node),
	})
}

func (node *AlterRoleStmt) UnmarshalJSON(input []byte) (err error) {
	var fields map[string]json.RawMessage

	err = json.Unmarshal(input, &fields)
	if err != nil {
		return
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

	if fields["action"] != nil {
		err = json.Unmarshal(fields["action"], &node.Action)
		if err != nil {
			return
		}
	}

	return
}
