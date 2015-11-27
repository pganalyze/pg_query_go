// Auto-generated from postgres/src/include/nodes/primnodes.h - DO NOT EDIT

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

	if fields["refarraytype"] != nil {
		err = json.Unmarshal(fields["refarraytype"], &node.Refarraytype)
		if err != nil {
			return
		}
	}

	if fields["refelemtype"] != nil {
		err = json.Unmarshal(fields["refelemtype"], &node.Refelemtype)
		if err != nil {
			return
		}
	}

	if fields["reftypmod"] != nil {
		err = json.Unmarshal(fields["reftypmod"], &node.Reftypmod)
		if err != nil {
			return
		}
	}

	if fields["refcollid"] != nil {
		err = json.Unmarshal(fields["refcollid"], &node.Refcollid)
		if err != nil {
			return
		}
	}

	if fields["refupperindexpr"] != nil {
		node.Refupperindexpr, err = UnmarshalNodeArrayJSON(fields["refupperindexpr"])
		if err != nil {
			return
		}
	}

	if fields["reflowerindexpr"] != nil {
		node.Reflowerindexpr, err = UnmarshalNodeArrayJSON(fields["reflowerindexpr"])
		if err != nil {
			return
		}
	}

	if fields["refexpr"] != nil {
		var nodePtr *Node
		nodePtr, err = UnmarshalNodePtrJSON(fields["refexpr"])
		if err != nil {
			return
		}
		if nodePtr != nil && *nodePtr != nil {
			val := (*nodePtr).(Expr)
			node.Refexpr = &val
		}
	}

	if fields["refassgnexpr"] != nil {
		var nodePtr *Node
		nodePtr, err = UnmarshalNodePtrJSON(fields["refassgnexpr"])
		if err != nil {
			return
		}
		if nodePtr != nil && *nodePtr != nil {
			val := (*nodePtr).(Expr)
			node.Refassgnexpr = &val
		}
	}

	return
}
