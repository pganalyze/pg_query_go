// Auto-generated - DO NOT EDIT

package pg_query

type RangeVar struct {
	Catalogname *string   `json:"catalogname"` /* the catalog (database) name, or NULL */
	Schemaname  *string   `json:"schemaname"`  /* the schema name, or NULL */
	Relname     *string   `json:"relname"`     /* the relation/sequence name */
	InhOpt      InhOption `json:"inhOpt"`      /* expand rel by inheritance? recursively act
	 * on children? */
	Relpersistence byte   `json:"relpersistence"` /* see RELPERSISTENCE_* in pg_class.h */
	Alias          *Alias `json:"alias"`          /* table alias & optional column aliases */
	Location       int    `json:"location"`       /* token location, or -1 if unknown */
}
