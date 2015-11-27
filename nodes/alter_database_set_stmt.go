// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

type AlterDatabaseSetStmt struct {
	Dbname  *string          `json:"dbname"`  /* database name */
	Setstmt *VariableSetStmt `json:"setstmt"` /* SET or RESET subcommand */
}

func (node AlterDatabaseSetStmt) MarshalJSON() ([]byte, error) {
	type AlterDatabaseSetStmtMarshalAlias AlterDatabaseSetStmt
	return json.Marshal(map[string]interface{}{
		"ALTERDATABASESETSTMT": (*AlterDatabaseSetStmtMarshalAlias)(&node),
	})
}

func (node *AlterDatabaseSetStmt) UnmarshalJSON(input []byte) (err error) {
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

	if fields["setstmt"] != nil {
		var nodePtr *Node
		nodePtr, err = UnmarshalNodePtrJSON(fields["setstmt"])
		if err != nil {
			return
		}
		if nodePtr != nil && *nodePtr != nil {
			val := (*nodePtr).(VariableSetStmt)
			node.Setstmt = &val
		}
	}

	return
}
