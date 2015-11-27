// Auto-generated from postgres/src/include/nodes/relation.h - DO NOT EDIT

package pg_query

import "encoding/json"

/*
 * PathKeys
 *
 * The sort ordering of a path is represented by a list of PathKey nodes.
 * An empty list implies no known ordering.  Otherwise the first item
 * represents the primary sort key, the second the first secondary sort key,
 * etc.  The value being sorted is represented by linking to an
 * EquivalenceClass containing that value and including pk_opfamily among its
 * ec_opfamilies.  The EquivalenceClass tells which collation to use, too.
 * This is a convenient method because it makes it trivial to detect
 * equivalent and closely-related orderings. (See optimizer/README for more
 * information.)
 *
 * Note: pk_strategy is either BTLessStrategyNumber (for ASC) or
 * BTGreaterStrategyNumber (for DESC).  We assume that all ordering-capable
 * index types will use btree-compatible strategy numbers.
 */
type PathKey struct {
	PkEclass     *EquivalenceClass `json:"pk_eclass"`      /* the value that is ordered */
	PkOpfamily   Oid               `json:"pk_opfamily"`    /* btree opfamily defining the ordering */
	PkStrategy   int               `json:"pk_strategy"`    /* sort direction (ASC or DESC) */
	PkNullsFirst bool              `json:"pk_nulls_first"` /* do NULLs come before normal values? */
}

func (node PathKey) MarshalJSON() ([]byte, error) {
	type PathKeyMarshalAlias PathKey
	return json.Marshal(map[string]interface{}{
		"PATHKEY": (*PathKeyMarshalAlias)(&node),
	})
}

func (node *PathKey) UnmarshalJSON(input []byte) (err error) {
	var fields map[string]json.RawMessage

	err = json.Unmarshal(input, &fields)
	if err != nil {
		return
	}

	if fields["pk_eclass"] != nil {
		var nodePtr *Node
		nodePtr, err = UnmarshalNodePtrJSON(fields["pk_eclass"])
		if err != nil {
			return
		}
		if nodePtr != nil && *nodePtr != nil {
			val := (*nodePtr).(EquivalenceClass)
			node.PkEclass = &val
		}
	}

	if fields["pk_opfamily"] != nil {
		err = json.Unmarshal(fields["pk_opfamily"], &node.PkOpfamily)
		if err != nil {
			return
		}
	}

	if fields["pk_strategy"] != nil {
		err = json.Unmarshal(fields["pk_strategy"], &node.PkStrategy)
		if err != nil {
			return
		}
	}

	if fields["pk_nulls_first"] != nil {
		err = json.Unmarshal(fields["pk_nulls_first"], &node.PkNullsFirst)
		if err != nil {
			return
		}
	}

	return
}
