// Auto-generated from postgres/src/include/nodes/primnodes.h - DO NOT EDIT

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
	var fields map[string]json.RawMessage

	err = json.Unmarshal(input, &fields)
	if err != nil {
		return
	}

	if fields["xpr"] != nil {
		err = json.Unmarshal(fields["xpr"], &node.Xpr)
		if err != nil {
			return
		}
	}

	if fields["paramkind"] != nil {
		err = json.Unmarshal(fields["paramkind"], &node.Paramkind)
		if err != nil {
			return
		}
	}

	if fields["paramid"] != nil {
		err = json.Unmarshal(fields["paramid"], &node.Paramid)
		if err != nil {
			return
		}
	}

	if fields["paramtype"] != nil {
		err = json.Unmarshal(fields["paramtype"], &node.Paramtype)
		if err != nil {
			return
		}
	}

	if fields["paramtypmod"] != nil {
		err = json.Unmarshal(fields["paramtypmod"], &node.Paramtypmod)
		if err != nil {
			return
		}
	}

	if fields["paramcollid"] != nil {
		err = json.Unmarshal(fields["paramcollid"], &node.Paramcollid)
		if err != nil {
			return
		}
	}

	if fields["location"] != nil {
		err = json.Unmarshal(fields["location"], &node.Location)
		if err != nil {
			return
		}
	}

	return
}
