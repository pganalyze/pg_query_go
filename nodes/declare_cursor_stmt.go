// Auto-generated - DO NOT EDIT

package pg_query

type DeclareCursorStmt struct {
	Portalname *string `json:"portalname"` /* name of the portal (cursor) */
	Options    int     `json:"options"`    /* bitmask of options (see above) */
	Query      Node    `json:"query"`      /* the raw SELECT query */
}
