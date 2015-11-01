// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

type CreateCastStmt struct {
	Sourcetype *TypeName       `json:"sourcetype"`
	Targettype *TypeName       `json:"targettype"`
	Func       *FuncWithArgs   `json:"func"`
	Context    CoercionContext `json:"context"`
	Inout      bool            `json:"inout"`
}

func (node CreateCastStmt) MarshalJSON() ([]byte, error) {
	type CreateCastStmtMarshalAlias CreateCastStmt
	return json.Marshal(map[string]interface{}{
		"CREATECASTSTMT": (*CreateCastStmtMarshalAlias)(&node),
	})
}

func (node *CreateCastStmt) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}
