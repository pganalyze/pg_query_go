// Auto-generated - DO NOT EDIT

package pg_query

type PathKey struct {
	PkEclass     *EquivalenceClass `json:"pk_eclass"`      /* the value that is ordered */
	PkOpfamily   Oid               `json:"pk_opfamily"`    /* btree opfamily defining the ordering */
	PkStrategy   int               `json:"pk_strategy"`    /* sort direction (ASC or DESC) */
	PkNullsFirst bool              `json:"pk_nulls_first"` /* do NULLs come before normal values? */
}
