// Auto-generated from postgres/src/include/nodes/relation.h - DO NOT EDIT

package pg_query

import "encoding/json"

/*
 * EquivalenceMember - one member expression of an EquivalenceClass
 *
 * em_is_child signifies that this element was built by transposing a member
 * for an appendrel parent relation to represent the corresponding expression
 * for an appendrel child.  These members are used for determining the
 * pathkeys of scans on the child relation and for explicitly sorting the
 * child when necessary to build a MergeAppend path for the whole appendrel
 * tree.  An em_is_child member has no impact on the properties of the EC as a
 * whole; in particular the EC's ec_relids field does NOT include the child
 * relation.  An em_is_child member should never be marked em_is_const nor
 * cause ec_has_const or ec_has_volatile to be set, either.  Thus, em_is_child
 * members are not really full-fledged members of the EC, but just reflections
 * or doppelgangers of real members.  Most operations on EquivalenceClasses
 * should ignore em_is_child members, and those that don't should test
 * em_relids to make sure they only consider relevant members.
 *
 * em_datatype is usually the same as exprType(em_expr), but can be
 * different when dealing with a binary-compatible opfamily; in particular
 * anyarray_ops would never work without this.  Use em_datatype when
 * looking up a specific btree operator to work with this expression.
 */
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
