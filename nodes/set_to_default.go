// Auto-generated - DO NOT EDIT

package pg_query

type SetToDefault struct {
	Xpr       Expr  `json:"xpr"`
	TypeId    Oid   `json:"typeId"`    /* type for substituted value */
	TypeMod   int32 `json:"typeMod"`   /* typemod for substituted value */
	Collation Oid   `json:"collation"` /* collation for the substituted value */
	Location  int   `json:"location"`  /* token location, or -1 if unknown */
}
