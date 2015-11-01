// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

type Const struct {
	Xpr         Expr  `json:"xpr"`
	Consttype   Oid   `json:"consttype"`   /* pg_type OID of the constant's datatype */
	Consttypmod int32 `json:"consttypmod"` /* typmod value, if any */
	Constcollid Oid   `json:"constcollid"` /* OID of collation, or InvalidOid if none */
	Constlen    int   `json:"constlen"`    /* typlen of the constant's datatype */
	Constvalue  Datum `json:"constvalue"`  /* the constant's value */
	Constisnull bool  `json:"constisnull"` /* whether the constant is null (if true,
	 * constvalue is undefined) */
	Constbyval bool `json:"constbyval"` /* whether this datatype is passed by value.
	 * If true, then all the information is stored
	 * in the Datum. If false, then the Datum
	 * contains a pointer to the information. */
	Location int `json:"location"` /* token location, or -1 if unknown */
}

func (node Const) MarshalJSON() ([]byte, error) {
	type ConstMarshalAlias Const
	return json.Marshal(map[string]interface{}{
		"CONST": (*ConstMarshalAlias)(&node),
	})
}

func (node *Const) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}
