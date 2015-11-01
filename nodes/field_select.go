// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

type FieldSelect struct {
	Xpr        Expr       `json:"xpr"`
	Arg        *Expr      `json:"arg"`        /* input expression */
	Fieldnum   AttrNumber `json:"fieldnum"`   /* attribute number of field to extract */
	Resulttype Oid        `json:"resulttype"` /* type of the field (result type of this
	 * node) */
	Resulttypmod int32 `json:"resulttypmod"` /* output typmod (usually -1) */
	Resultcollid Oid   `json:"resultcollid"` /* OID of collation of the field */
}

func (node FieldSelect) MarshalJSON() ([]byte, error) {
	type FieldSelectMarshalAlias FieldSelect
	return json.Marshal(map[string]interface{}{
		"FIELDSELECT": (*FieldSelectMarshalAlias)(&node),
	})
}

func (node *FieldSelect) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}
