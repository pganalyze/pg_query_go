// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

type CreatedbStmt struct {
	Dbname  *string `json:"dbname"`  /* name of database to create */
	Options []Node  `json:"options"` /* List of DefElem nodes */
}

func (node CreatedbStmt) MarshalJSON() ([]byte, error) {
	type CreatedbStmtMarshalAlias CreatedbStmt
	return json.Marshal(map[string]interface{}{
		"CREATEDBSTMT": (*CreatedbStmtMarshalAlias)(&node),
	})
}

func (node *CreatedbStmt) UnmarshalJSON(input []byte) (err error) {
	var fields map[string]json.RawMessage

	err = json.Unmarshal(input, &fields)
	if err != nil {
		return
	}

	if fields["dbname"] != nil {
		err = json.Unmarshal(fields["dbname"], &node.Dbname)
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
