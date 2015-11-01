// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

type ViewStmt struct {
	View            *RangeVar       `json:"view"`            /* the view to be created */
	Aliases         []Node          `json:"aliases"`         /* target column names */
	Query           Node            `json:"query"`           /* the SELECT query */
	Replace         bool            `json:"replace"`         /* replace an existing view? */
	Options         []Node          `json:"options"`         /* options from WITH clause */
	WithCheckOption ViewCheckOption `json:"withCheckOption"` /* WITH CHECK OPTION */
}

func (node ViewStmt) MarshalJSON() ([]byte, error) {
	type ViewStmtMarshalAlias ViewStmt
	return json.Marshal(map[string]interface{}{
		"VIEWSTMT": (*ViewStmtMarshalAlias)(&node),
	})
}

func (node *ViewStmt) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}
