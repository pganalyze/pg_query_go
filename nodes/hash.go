// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

type Hash struct {
	Plan          Plan       `json:"plan"`
	SkewTable     Oid        `json:"skewTable"`     /* outer join key's table OID, or InvalidOid */
	SkewColumn    AttrNumber `json:"skewColumn"`    /* outer join key's column #, or zero */
	SkewInherit   bool       `json:"skewInherit"`   /* is outer join rel an inheritance tree? */
	SkewColType   Oid        `json:"skewColType"`   /* datatype of the outer key column */
	SkewColTypmod int32      `json:"skewColTypmod"` /* typmod of the outer key column */
	/* all other info is in the parent HashJoin node */
}

func (node Hash) MarshalJSON() ([]byte, error) {
	type HashMarshalAlias Hash
	return json.Marshal(map[string]interface{}{
		"HASH": (*HashMarshalAlias)(&node),
	})
}

func (node *Hash) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}
