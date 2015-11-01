// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

type PrepareStmt struct {
	Name     *string `json:"name"`     /* Name of plan, arbitrary */
	Argtypes []Node  `json:"argtypes"` /* Types of parameters (List of TypeName) */
	Query    Node    `json:"query"`    /* The query itself (as a raw parsetree) */
}

func (node PrepareStmt) MarshalJSON() ([]byte, error) {
	type PrepareStmtMarshalAlias PrepareStmt
	return json.Marshal(map[string]interface{}{
		"PREPARESTMT": (*PrepareStmtMarshalAlias)(&node),
	})
}

func (node *PrepareStmt) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}
