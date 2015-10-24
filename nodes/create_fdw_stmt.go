// Auto-generated - DO NOT EDIT

package pg_query

type CreateFdwStmt struct {
	Fdwname     *string `json:"fdwname"`      /* foreign-data wrapper name */
	FuncOptions []Node  `json:"func_options"` /* HANDLER/VALIDATOR options */
	Options     []Node  `json:"options"`      /* generic options to FDW */
}
