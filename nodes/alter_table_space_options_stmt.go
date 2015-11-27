// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

type AlterTableSpaceOptionsStmt struct {
	Tablespacename *string `json:"tablespacename"`
	Options        []Node  `json:"options"`
	IsReset        bool    `json:"isReset"`
}

func (node AlterTableSpaceOptionsStmt) MarshalJSON() ([]byte, error) {
	type AlterTableSpaceOptionsStmtMarshalAlias AlterTableSpaceOptionsStmt
	return json.Marshal(map[string]interface{}{
		"ALTERTABLESPACEOPTIONSSTMT": (*AlterTableSpaceOptionsStmtMarshalAlias)(&node),
	})
}

func (node *AlterTableSpaceOptionsStmt) UnmarshalJSON(input []byte) (err error) {
	var fields map[string]json.RawMessage

	err = json.Unmarshal(input, &fields)
	if err != nil {
		return
	}

	if fields["tablespacename"] != nil {
		err = json.Unmarshal(fields["tablespacename"], &node.Tablespacename)
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

	if fields["isReset"] != nil {
		err = json.Unmarshal(fields["isReset"], &node.IsReset)
		if err != nil {
			return
		}
	}

	return
}
