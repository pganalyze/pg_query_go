// Auto-generated from postgres/src/include/nodes/primnodes.h - DO NOT EDIT

package pg_query

import "encoding/json"

type CoerceViaIO struct {
	Xpr        Expr  `json:"xpr"`
	Arg        *Expr `json:"arg"`        /* input expression */
	Resulttype Oid   `json:"resulttype"` /* output type of coercion */

	/* output typmod is not stored, but is presumed -1 */
	Resultcollid Oid          `json:"resultcollid"` /* OID of collation, or InvalidOid if none */
	Coerceformat CoercionForm `json:"coerceformat"` /* how to display this node */
	Location     int          `json:"location"`     /* token location, or -1 if unknown */
}

func (node CoerceViaIO) MarshalJSON() ([]byte, error) {
	type CoerceViaIOMarshalAlias CoerceViaIO
	return json.Marshal(map[string]interface{}{
		"COERCEVIAIO": (*CoerceViaIOMarshalAlias)(&node),
	})
}

func (node *CoerceViaIO) UnmarshalJSON(input []byte) (err error) {
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

	if fields["arg"] != nil {
		var nodePtr *Node
		nodePtr, err = UnmarshalNodePtrJSON(fields["arg"])
		if err != nil {
			return
		}
		if nodePtr != nil && *nodePtr != nil {
			val := (*nodePtr).(Expr)
			node.Arg = &val
		}
	}

	if fields["resulttype"] != nil {
		err = json.Unmarshal(fields["resulttype"], &node.Resulttype)
		if err != nil {
			return
		}
	}

	if fields["resultcollid"] != nil {
		err = json.Unmarshal(fields["resultcollid"], &node.Resultcollid)
		if err != nil {
			return
		}
	}

	if fields["coerceformat"] != nil {
		err = json.Unmarshal(fields["coerceformat"], &node.Coerceformat)
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
