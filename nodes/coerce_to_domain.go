// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

type CoerceToDomain struct {
	Xpr            Expr         `json:"xpr"`
	Arg            *Expr        `json:"arg"`            /* input expression */
	Resulttype     Oid          `json:"resulttype"`     /* domain type ID (result type) */
	Resulttypmod   int32        `json:"resulttypmod"`   /* output typmod (currently always -1) */
	Resultcollid   Oid          `json:"resultcollid"`   /* OID of collation, or InvalidOid if none */
	Coercionformat CoercionForm `json:"coercionformat"` /* how to display this node */
	Location       int          `json:"location"`       /* token location, or -1 if unknown */
}

func (node CoerceToDomain) MarshalJSON() ([]byte, error) {
	type CoerceToDomainMarshalAlias CoerceToDomain
	return json.Marshal(map[string]interface{}{
		"COERCETODOMAIN": (*CoerceToDomainMarshalAlias)(&node),
	})
}

func (node *CoerceToDomain) UnmarshalJSON(input []byte) (err error) {
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

	if fields["resulttypmod"] != nil {
		err = json.Unmarshal(fields["resulttypmod"], &node.Resulttypmod)
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

	if fields["coercionformat"] != nil {
		err = json.Unmarshal(fields["coercionformat"], &node.Coercionformat)
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
