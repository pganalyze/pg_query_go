// Auto-generated from postgres/src/include/nodes/relation.h - DO NOT EDIT

package pg_query

import "encoding/json"

type EquivalenceMember struct {
	EmExpr           *Expr    `json:"em_expr"`            /* the expression represented */
	EmRelids         []uint32 `json:"em_relids"`          /* all relids appearing in em_expr */
	EmNullableRelids []uint32 `json:"em_nullable_relids"` /* nullable by lower outer joins */
	EmIsConst        bool     `json:"em_is_const"`        /* expression is pseudoconstant? */
	EmIsChild        bool     `json:"em_is_child"`        /* derived version for a child relation? */
	EmDatatype       Oid      `json:"em_datatype"`        /* the "nominal type" used by the opfamily */
}

func (node EquivalenceMember) MarshalJSON() ([]byte, error) {
	type EquivalenceMemberMarshalAlias EquivalenceMember
	return json.Marshal(map[string]interface{}{
		"EQUIVALENCEMEMBER": (*EquivalenceMemberMarshalAlias)(&node),
	})
}

func (node *EquivalenceMember) UnmarshalJSON(input []byte) (err error) {
	var fields map[string]json.RawMessage

	err = json.Unmarshal(input, &fields)
	if err != nil {
		return
	}

	if fields["em_expr"] != nil {
		var nodePtr *Node
		nodePtr, err = UnmarshalNodePtrJSON(fields["em_expr"])
		if err != nil {
			return
		}
		if nodePtr != nil && *nodePtr != nil {
			val := (*nodePtr).(Expr)
			node.EmExpr = &val
		}
	}

	if fields["em_relids"] != nil {
		err = json.Unmarshal(fields["em_relids"], &node.EmRelids)
		if err != nil {
			return
		}
	}

	if fields["em_nullable_relids"] != nil {
		err = json.Unmarshal(fields["em_nullable_relids"], &node.EmNullableRelids)
		if err != nil {
			return
		}
	}

	if fields["em_is_const"] != nil {
		err = json.Unmarshal(fields["em_is_const"], &node.EmIsConst)
		if err != nil {
			return
		}
	}

	if fields["em_is_child"] != nil {
		err = json.Unmarshal(fields["em_is_child"], &node.EmIsChild)
		if err != nil {
			return
		}
	}

	if fields["em_datatype"] != nil {
		err = json.Unmarshal(fields["em_datatype"], &node.EmDatatype)
		if err != nil {
			return
		}
	}

	return
}
