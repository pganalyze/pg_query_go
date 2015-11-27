// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

type AlterSystemStmt struct {
	Setstmt *VariableSetStmt `json:"setstmt"` /* SET subcommand */
}

func (node AlterSystemStmt) MarshalJSON() ([]byte, error) {
	type AlterSystemStmtMarshalAlias AlterSystemStmt
	return json.Marshal(map[string]interface{}{
		"ALTERSYSTEMSTMT": (*AlterSystemStmtMarshalAlias)(&node),
	})
}

func (node *AlterSystemStmt) UnmarshalJSON(input []byte) (err error) {
	var fields map[string]json.RawMessage

	err = json.Unmarshal(input, &fields)
	if err != nil {
		return
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
