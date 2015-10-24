// Auto-generated - DO NOT EDIT

package pg_query

type CreateExtensionStmt struct {
	Extname     *string `json:"extname"`
	IfNotExists bool    `json:"if_not_exists"` /* just do nothing if it already exists? */
	Options     []Node  `json:"options"`       /* List of DefElem nodes */
}
