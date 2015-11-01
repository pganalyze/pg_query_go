// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

type ArrayRef struct {
	Xpr             Expr   `json:"xpr"`
	Refarraytype    Oid    `json:"refarraytype"`    /* type of the array proper */
	Refelemtype     Oid    `json:"refelemtype"`     /* type of the array elements */
	Reftypmod       int32  `json:"reftypmod"`       /* typmod of the array (and elements too) */
	Refcollid       Oid    `json:"refcollid"`       /* OID of collation, or InvalidOid if none */
	Refupperindexpr []Node `json:"refupperindexpr"` /* expressions that evaluate to upper array
	 * indexes */
	Reflowerindexpr []Node `json:"reflowerindexpr"` /* expressions that evaluate to lower array
	 * indexes */
	Refexpr *Expr `json:"refexpr"` /* the expression that evaluates to an array
	 * value */
	Refassgnexpr *Expr `json:"refassgnexpr"` /* expression for the source value, or NULL if
	 * fetch */
}

func (node ArrayRef) MarshalJSON() ([]byte, error) {
	type ArrayRefMarshalAlias ArrayRef
	return json.Marshal(map[string]interface{}{
		"ARRAYREF": (*ArrayRefMarshalAlias)(&node),
	})
}

func (node *ArrayRef) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}
