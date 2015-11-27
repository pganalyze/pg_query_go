// Auto-generated from postgres/src/include/nodes/relation.h - DO NOT EDIT

package pg_query

import "encoding/json"

/*
 * EquivalenceClasses
 *
 * Whenever we can determine that a mergejoinable equality clause A = B is
 * not delayed by any outer join, we create an EquivalenceClass containing
 * the expressions A and B to record this knowledge.  If we later find another
 * equivalence B = C, we add C to the existing EquivalenceClass; this may
 * require merging two existing EquivalenceClasses.  At the end of the qual
 * distribution process, we have sets of values that are known all transitively
 * equal to each other, where "equal" is according to the rules of the btree
 * operator family(s) shown in ec_opfamilies, as well as the collation shown
 * by ec_collation.  (We restrict an EC to contain only equalities whose
 * operators belong to the same set of opfamilies.  This could probably be
 * relaxed, but for now it's not worth the trouble, since nearly all equality
 * operators belong to only one btree opclass anyway.  Similarly, we suppose
 * that all or none of the input datatypes are collatable, so that a single
 * collation value is sufficient.)
 *
 * We also use EquivalenceClasses as the base structure for PathKeys, letting
 * us represent knowledge about different sort orderings being equivalent.
 * Since every PathKey must reference an EquivalenceClass, we will end up
 * with single-member EquivalenceClasses whenever a sort key expression has
 * not been equivalenced to anything else.  It is also possible that such an
 * EquivalenceClass will contain a volatile expression ("ORDER BY random()"),
 * which is a case that can't arise otherwise since clauses containing
 * volatile functions are never considered mergejoinable.  We mark such
 * EquivalenceClasses specially to prevent them from being merged with
 * ordinary EquivalenceClasses.  Also, for volatile expressions we have
 * to be careful to match the EquivalenceClass to the correct targetlist
 * entry: consider SELECT random() AS a, random() AS b ... ORDER BY b,a.
 * So we record the SortGroupRef of the originating sort clause.
 *
 * We allow equality clauses appearing below the nullable side of an outer join
 * to form EquivalenceClasses, but these have a slightly different meaning:
 * the included values might be all NULL rather than all the same non-null
 * values.  See src/backend/optimizer/README for more on that point.
 *
 * NB: if ec_merged isn't NULL, this class has been merged into another, and
 * should be ignored in favor of using the pointed-to class.
 */
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
