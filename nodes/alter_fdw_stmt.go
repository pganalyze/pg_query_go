// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

type AlterFdwStmt struct {
	Fdwname     *string `json:"fdwname"`      /* foreign-data wrapper name */
	FuncOptions []Node  `json:"func_options"` /* HANDLER/VALIDATOR options */
	Options     []Node  `json:"options"`      /* generic options to FDW */
}

func (node AlterFdwStmt) MarshalJSON() ([]byte, error) {
	type AlterFdwStmtMarshalAlias AlterFdwStmt
	return json.Marshal(map[string]interface{}{
		"ALTERFDWSTMT": (*AlterFdwStmtMarshalAlias)(&node),
	})
}

func (node *AlterFdwStmt) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}
