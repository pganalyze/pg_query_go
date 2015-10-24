// Auto-generated - DO NOT EDIT

package pg_query

type ArrayExpr struct {
	Xpr           Expr   `json:"xpr"`
	ArrayTypeid   Oid    `json:"array_typeid"`   /* type of expression result */
	ArrayCollid   Oid    `json:"array_collid"`   /* OID of collation, or InvalidOid if none */
	ElementTypeid Oid    `json:"element_typeid"` /* common type of array elements */
	Elements      []Node `json:"elements"`       /* the array elements or sub-arrays */
	Multidims     bool   `json:"multidims"`      /* true if elements are sub-arrays */
	Location      int    `json:"location"`       /* token location, or -1 if unknown */
}
