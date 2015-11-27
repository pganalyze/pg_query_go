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
	var fields map[string]json.RawMessage

	err = json.Unmarshal(input, &fields)
	if err != nil {
		return
	}

	if fields["typeName"] != nil {
		node.TypeName, err = UnmarshalNodeArrayJSON(fields["typeName"])
		if err != nil {
			return
		}
	}

	if fields["vals"] != nil {
		node.Vals, err = UnmarshalNodeArrayJSON(fields["vals"])
		if err != nil {
			return
		}
	}

	return
}
