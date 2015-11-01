// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

type CreateFdwStmt struct {
	Fdwname     *string `json:"fdwname"`      /* foreign-data wrapper name */
	FuncOptions []Node  `json:"func_options"` /* HANDLER/VALIDATOR options */
	Options     []Node  `json:"options"`      /* generic options to FDW */
}

func (node CreateFdwStmt) MarshalJSON() ([]byte, error) {
	type CreateFdwStmtMarshalAlias CreateFdwStmt
	return json.Marshal(map[string]interface{}{
		"CREATEFDWSTMT": (*CreateFdwStmtMarshalAlias)(&node),
	})
}

func (node *CreateFdwStmt) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}
