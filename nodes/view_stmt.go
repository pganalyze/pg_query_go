// Auto-generated - DO NOT EDIT

package pg_query

type ViewStmt struct {
	View            *RangeVar       `json:"view"`            /* the view to be created */
	Aliases         []Node          `json:"aliases"`         /* target column names */
	Query           Node            `json:"query"`           /* the SELECT query */
	Replace         bool            `json:"replace"`         /* replace an existing view? */
	Options         []Node          `json:"options"`         /* options from WITH clause */
	WithCheckOption ViewCheckOption `json:"withCheckOption"` /* WITH CHECK OPTION */
}
