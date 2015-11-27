// Auto-generated from postgres/src/include/nodes/relation.h - DO NOT EDIT

package pg_query

import "encoding/json"

type EquivalenceClass struct {
	EcOpfamilies []Node   `json:"ec_opfamilies"` /* btree operator family OIDs */
	EcCollation  Oid      `json:"ec_collation"`  /* collation, if datatypes are collatable */
	EcMembers    []Node   `json:"ec_members"`    /* list of EquivalenceMembers */
	EcSources    []Node   `json:"ec_sources"`    /* list of generating RestrictInfos */
	EcDerives    []Node   `json:"ec_derives"`    /* list of derived RestrictInfos */
	EcRelids     []uint32 `json:"ec_relids"`     /* all relids appearing in ec_members, except
	 * for child members (see below) */

	EcHasConst       bool              `json:"ec_has_const"`        /* any pseudoconstants in ec_members? */
	EcHasVolatile    bool              `json:"ec_has_volatile"`     /* the (sole) member is a volatile expr */
	EcBelowOuterJoin bool              `json:"ec_below_outer_join"` /* equivalence applies below an OJ */
	EcBroken         bool              `json:"ec_broken"`           /* failed to generate needed clauses? */
	EcSortref        Index             `json:"ec_sortref"`          /* originating sortclause label, or 0 */
	EcMerged         *EquivalenceClass `json:"ec_merged"`           /* set if merged into another EC */
}

func (node EquivalenceClass) MarshalJSON() ([]byte, error) {
	type EquivalenceClassMarshalAlias EquivalenceClass
	return json.Marshal(map[string]interface{}{
		"EQUIVALENCECLASS": (*EquivalenceClassMarshalAlias)(&node),
	})
}

func (node *EquivalenceClass) UnmarshalJSON(input []byte) (err error) {
	var fields map[string]json.RawMessage

	err = json.Unmarshal(input, &fields)
	if err != nil {
		return
	}

	if fields["ec_opfamilies"] != nil {
		node.EcOpfamilies, err = UnmarshalNodeArrayJSON(fields["ec_opfamilies"])
		if err != nil {
			return
		}
	}

	if fields["ec_collation"] != nil {
		err = json.Unmarshal(fields["ec_collation"], &node.EcCollation)
		if err != nil {
			return
		}
	}

	if fields["ec_members"] != nil {
		node.EcMembers, err = UnmarshalNodeArrayJSON(fields["ec_members"])
		if err != nil {
			return
		}
	}

	if fields["ec_sources"] != nil {
		node.EcSources, err = UnmarshalNodeArrayJSON(fields["ec_sources"])
		if err != nil {
			return
		}
	}

	if fields["ec_derives"] != nil {
		node.EcDerives, err = UnmarshalNodeArrayJSON(fields["ec_derives"])
		if err != nil {
			return
		}
	}

	if fields["ec_relids"] != nil {
		err = json.Unmarshal(fields["ec_relids"], &node.EcRelids)
		if err != nil {
			return
		}
	}

	if fields["ec_has_const"] != nil {
		err = json.Unmarshal(fields["ec_has_const"], &node.EcHasConst)
		if err != nil {
			return
		}
	}

	if fields["ec_has_volatile"] != nil {
		err = json.Unmarshal(fields["ec_has_volatile"], &node.EcHasVolatile)
		if err != nil {
			return
		}
	}

	if fields["ec_below_outer_join"] != nil {
		err = json.Unmarshal(fields["ec_below_outer_join"], &node.EcBelowOuterJoin)
		if err != nil {
			return
		}
	}

	if fields["ec_broken"] != nil {
		err = json.Unmarshal(fields["ec_broken"], &node.EcBroken)
		if err != nil {
			return
		}
	}

	if fields["ec_sortref"] != nil {
		err = json.Unmarshal(fields["ec_sortref"], &node.EcSortref)
		if err != nil {
			return
		}
	}

	if fields["ec_merged"] != nil {
		var nodePtr *Node
		nodePtr, err = UnmarshalNodePtrJSON(fields["ec_merged"])
		if err != nil {
			return
		}
		if nodePtr != nil && *nodePtr != nil {
			val := (*nodePtr).(EquivalenceClass)
			node.EcMerged = &val
		}
	}

	return
}
