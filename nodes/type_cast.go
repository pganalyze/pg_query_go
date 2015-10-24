// Auto-generated - DO NOT EDIT

package pg_query

type TypeCast struct {
	Arg      Node      `json:"arg"`      /* the expression being casted */
	TypeName *TypeName `json:"typeName"` /* the target type */
	Location int       `json:"location"` /* token location, or -1 if unknown */
}
