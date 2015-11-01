// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

type Param struct {
	Xpr         Expr      `json:"xpr"`
	Paramkind   ParamKind `json:"paramkind"`   /* kind of parameter. See above */
	Paramid     int       `json:"paramid"`     /* numeric ID for parameter */
	Paramtype   Oid       `json:"paramtype"`   /* pg_type OID of parameter's datatype */
	Paramtypmod int32     `json:"paramtypmod"` /* typmod value, if known */
	Paramcollid Oid       `json:"paramcollid"` /* OID of collation, or InvalidOid if none */
	Location    int       `json:"location"`    /* token location, or -1 if unknown */
}

func (node Param) MarshalJSON() ([]byte, error) {
	type ParamMarshalAlias Param
	return json.Marshal(map[string]interface{}{
		"PARAM": (*ParamMarshalAlias)(&node),
	})
}

func (node *Param) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}
