package pg_query

type Alias struct {
	Aliasname *string `json:"aliasname"` /* aliased rel name (never qualified) */
	Colnames  []Node  `json:"colnames"`  /* optional list of column aliases */
}
