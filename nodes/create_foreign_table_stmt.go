// Auto-generated from postgres/src/include/nodes/parsenodes.h - DO NOT EDIT

package pg_query

import "encoding/json"

type CreateForeignTableStmt struct {
	Base       CreateStmt `json:"base"`
	Servername *string    `json:"servername"`
	Options    []Node     `json:"options"`
}

func (node CreateForeignTableStmt) MarshalJSON() ([]byte, error) {
	type CreateForeignTableStmtMarshalAlias CreateForeignTableStmt
	return json.Marshal(map[string]interface{}{
		"CREATEFOREIGNTABLESTMT": (*CreateForeignTableStmtMarshalAlias)(&node),
	})
}

func (node *CreateForeignTableStmt) UnmarshalJSON(input []byte) (err error) {
	var fields map[string]json.RawMessage

	err = json.Unmarshal(input, &fields)
	if err != nil {
		return
	}

	if fields["base"] != nil {
		err = json.Unmarshal(fields["base"], &node.Base)
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
