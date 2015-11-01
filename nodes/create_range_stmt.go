// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

type CreateRangeStmt struct {
	TypeName []Node `json:"typeName"` /* qualified name (list of Value strings) */
	Params   []Node `json:"params"`   /* range parameters (list of DefElem) */
}

func (node CreateRangeStmt) MarshalJSON() ([]byte, error) {
	type CreateRangeStmtMarshalAlias CreateRangeStmt
	return json.Marshal(map[string]interface{}{
		"CREATERANGESTMT": (*CreateRangeStmtMarshalAlias)(&node),
	})
}

func (node *CreateRangeStmt) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}
