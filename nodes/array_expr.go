// Auto-generated from postgres/src/include/nodes/primnodes.h - DO NOT EDIT

package pg_query

import "encoding/json"

type ArrayExpr struct {
	Xpr           Expr   `json:"xpr"`
	ArrayTypeid   Oid    `json:"array_typeid"`   /* type of expression result */
	ArrayCollid   Oid    `json:"array_collid"`   /* OID of collation, or InvalidOid if none */
	ElementTypeid Oid    `json:"element_typeid"` /* common type of array elements */
	Elements      []Node `json:"elements"`       /* the array elements or sub-arrays */
	Multidims     bool   `json:"multidims"`      /* true if elements are sub-arrays */
	Location      int    `json:"location"`       /* token location, or -1 if unknown */
}

func (node ArrayExpr) MarshalJSON() ([]byte, error) {
	type ArrayExprMarshalAlias ArrayExpr
	return json.Marshal(map[string]interface{}{
		"ARRAY": (*ArrayExprMarshalAlias)(&node),
	})
}

func (node *ArrayExpr) UnmarshalJSON(input []byte) (err error) {
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

	if fields["array_typeid"] != nil {
		err = json.Unmarshal(fields["array_typeid"], &node.ArrayTypeid)
		if err != nil {
			return
		}
	}

	if fields["array_collid"] != nil {
		err = json.Unmarshal(fields["array_collid"], &node.ArrayCollid)
		if err != nil {
			return
		}
	}

	if fields["element_typeid"] != nil {
		err = json.Unmarshal(fields["element_typeid"], &node.ElementTypeid)
		if err != nil {
			return
		}
	}

	if fields["elements"] != nil {
		node.Elements, err = UnmarshalNodeArrayJSON(fields["elements"])
		if err != nil {
			return
		}
	}

	if fields["multidims"] != nil {
		err = json.Unmarshal(fields["multidims"], &node.Multidims)
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
