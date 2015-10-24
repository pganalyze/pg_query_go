// Auto-generated - DO NOT EDIT

package pg_query

type FetchStmt struct {
	Direction  FetchDirection `json:"direction"`  /* see above */
	HowMany    int64          `json:"howMany"`    /* number of rows, or position argument */
	Portalname *string        `json:"portalname"` /* name of portal (cursor) */
	Ismove     bool           `json:"ismove"`     /* TRUE if MOVE */
}
