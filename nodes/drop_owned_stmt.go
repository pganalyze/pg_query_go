// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

type DropOwnedStmt struct {
	Roles    []Node       `json:"roles"`
	Behavior DropBehavior `json:"behavior"`
}

func (node DropOwnedStmt) MarshalJSON() ([]byte, error) {
	type DropOwnedStmtMarshalAlias DropOwnedStmt
	return json.Marshal(map[string]interface{}{
		"DROPOWNEDSTMT": (*DropOwnedStmtMarshalAlias)(&node),
	})
}

func (node *DropOwnedStmt) UnmarshalJSON(input []byte) (err error) {
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

	if fields["behavior"] != nil {
		err = json.Unmarshal(fields["behavior"], &node.Behavior)
		if err != nil {
			return
		}
	}

	return
}
