// Auto-generated - DO NOT EDIT

package pg_query

type TargetEntry struct {
	Xpr             Expr       `json:"xpr"`
	Expr            *Expr      `json:"expr"`            /* expression to evaluate */
	Resno           AttrNumber `json:"resno"`           /* attribute number (see notes above) */
	Resname         *string    `json:"resname"`         /* name of the column (could be NULL) */
	Ressortgroupref Index      `json:"ressortgroupref"` /* nonzero if referenced by a sort/group
	 * clause */
	Resorigtbl Oid        `json:"resorigtbl"` /* OID of column's source table */
	Resorigcol AttrNumber `json:"resorigcol"` /* column's number in source table */
	Resjunk    bool       `json:"resjunk"`    /* set to true to eliminate the attribute from
	 * final target list */
}
