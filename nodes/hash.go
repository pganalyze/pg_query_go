// Auto-generated from postgres/src/include/nodes/plannodes.h - DO NOT EDIT

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
	var fields map[string]json.RawMessage

	err = json.Unmarshal(input, &fields)
	if err != nil {
		return
	}

	if fields["plan"] != nil {
		err = json.Unmarshal(fields["plan"], &node.Plan)
		if err != nil {
			return
		}
	}

	if fields["skewTable"] != nil {
		err = json.Unmarshal(fields["skewTable"], &node.SkewTable)
		if err != nil {
			return
		}
	}

	if fields["skewColumn"] != nil {
		err = json.Unmarshal(fields["skewColumn"], &node.SkewColumn)
		if err != nil {
			return
		}
	}

	if fields["skewInherit"] != nil {
		err = json.Unmarshal(fields["skewInherit"], &node.SkewInherit)
		if err != nil {
			return
		}
	}

	if fields["skewColType"] != nil {
		err = json.Unmarshal(fields["skewColType"], &node.SkewColType)
		if err != nil {
			return
		}
	}

	if fields["skewColTypmod"] != nil {
		err = json.Unmarshal(fields["skewColTypmod"], &node.SkewColTypmod)
		if err != nil {
			return
		}
	}

	return
}
