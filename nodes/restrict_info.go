// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

type RestrictInfo struct {
	Clause *Expr `json:"clause"` /* the represented clause of WHERE or JOIN */

	IsPushedDown bool `json:"is_pushed_down"` /* TRUE if clause was pushed down in level */

	OuterjoinDelayed bool `json:"outerjoin_delayed"` /* TRUE if delayed by lower outer join */

	CanJoin bool `json:"can_join"` /* see comment above */

	Pseudoconstant bool `json:"pseudoconstant"` /* see comment above */

	/* The set of relids (varnos) actually referenced in the clause: */
	ClauseRelids []uint32 `json:"clause_relids"`

	/* The set of relids required to evaluate the clause: */
	RequiredRelids []uint32 `json:"required_relids"`

	/* If an outer-join clause, the outer-side relations, else NULL: */
	OuterRelids []uint32 `json:"outer_relids"`

	/* The relids used in the clause that are nullable by lower outer joins: */
	NullableRelids []uint32 `json:"nullable_relids"`

	/* These fields are set for any binary opclause: */
	LeftRelids  []uint32 `json:"left_relids"`  /* relids in left side of clause */
	RightRelids []uint32 `json:"right_relids"` /* relids in right side of clause */

	/* This field is NULL unless clause is an OR clause: */
	Orclause *Expr `json:"orclause"` /* modified clause with RestrictInfos */

	/* This field is NULL unless clause is potentially redundant: */
	ParentEc *EquivalenceClass `json:"parent_ec"` /* generating EquivalenceClass */

	/* cache space for cost and selectivity */
	EvalCost  QualCost    `json:"eval_cost"`  /* eval cost of clause; -1 if not yet set */
	NormSelec Selectivity `json:"norm_selec"` /* selectivity for "normal" (JOIN_INNER)
	 * semantics; -1 if not yet set; >1 means a
	 * redundant clause */

	OuterSelec Selectivity `json:"outer_selec"` /* selectivity for outer join semantics; -1 if
	 * not yet set */

	/* valid if clause is mergejoinable, else NIL */
	Mergeopfamilies []Node `json:"mergeopfamilies"` /* opfamilies containing clause operator */

	/* cache space for mergeclause processing; NULL if not yet set */
	LeftEc       *EquivalenceClass  `json:"left_ec"`       /* EquivalenceClass containing lefthand */
	RightEc      *EquivalenceClass  `json:"right_ec"`      /* EquivalenceClass containing righthand */
	LeftEm       *EquivalenceMember `json:"left_em"`       /* EquivalenceMember for lefthand */
	RightEm      *EquivalenceMember `json:"right_em"`      /* EquivalenceMember for righthand */
	ScanselCache []Node             `json:"scansel_cache"` /* list of MergeScanSelCache structs */

	/* transient workspace for use while considering a specific join path */
	OuterIsLeft bool `json:"outer_is_left"` /* T = outer var on left, F = on right */

	/* valid if clause is hashjoinable, else InvalidOid: */
	Hashjoinoperator Oid `json:"hashjoinoperator"` /* copy of clause operator */

	/* cache space for hashclause processing; -1 if not yet set */
	LeftBucketsize  Selectivity `json:"left_bucketsize"`  /* avg bucketsize of left side */
	RightBucketsize Selectivity `json:"right_bucketsize"` /* avg bucketsize of right side */
}

func (node RestrictInfo) MarshalJSON() ([]byte, error) {
	type RestrictInfoMarshalAlias RestrictInfo
	return json.Marshal(map[string]interface{}{
		"RESTRICTINFO": (*RestrictInfoMarshalAlias)(&node),
	})
}

func (node *RestrictInfo) UnmarshalJSON(input []byte) (err error) {
	var fields map[string]json.RawMessage

	err = json.Unmarshal(input, &fields)
	if err != nil {
		return
	}

	if fields["clause"] != nil {
		var nodePtr *Node
		nodePtr, err = UnmarshalNodePtrJSON(fields["clause"])
		if err != nil {
			return
		}
		if nodePtr != nil && *nodePtr != nil {
			val := (*nodePtr).(Expr)
			node.Clause = &val
		}
	}

	if fields["is_pushed_down"] != nil {
		err = json.Unmarshal(fields["is_pushed_down"], &node.IsPushedDown)
		if err != nil {
			return
		}
	}

	if fields["outerjoin_delayed"] != nil {
		err = json.Unmarshal(fields["outerjoin_delayed"], &node.OuterjoinDelayed)
		if err != nil {
			return
		}
	}

	if fields["can_join"] != nil {
		err = json.Unmarshal(fields["can_join"], &node.CanJoin)
		if err != nil {
			return
		}
	}

	if fields["pseudoconstant"] != nil {
		err = json.Unmarshal(fields["pseudoconstant"], &node.Pseudoconstant)
		if err != nil {
			return
		}
	}

	if fields["clause_relids"] != nil {
		err = json.Unmarshal(fields["clause_relids"], &node.ClauseRelids)
		if err != nil {
			return
		}
	}

	if fields["required_relids"] != nil {
		err = json.Unmarshal(fields["required_relids"], &node.RequiredRelids)
		if err != nil {
			return
		}
	}

	if fields["outer_relids"] != nil {
		err = json.Unmarshal(fields["outer_relids"], &node.OuterRelids)
		if err != nil {
			return
		}
	}

	if fields["nullable_relids"] != nil {
		err = json.Unmarshal(fields["nullable_relids"], &node.NullableRelids)
		if err != nil {
			return
		}
	}

	if fields["left_relids"] != nil {
		err = json.Unmarshal(fields["left_relids"], &node.LeftRelids)
		if err != nil {
			return
		}
	}

	if fields["right_relids"] != nil {
		err = json.Unmarshal(fields["right_relids"], &node.RightRelids)
		if err != nil {
			return
		}
	}

	if fields["orclause"] != nil {
		var nodePtr *Node
		nodePtr, err = UnmarshalNodePtrJSON(fields["orclause"])
		if err != nil {
			return
		}
		if nodePtr != nil && *nodePtr != nil {
			val := (*nodePtr).(Expr)
			node.Orclause = &val
		}
	}

	if fields["parent_ec"] != nil {
		var nodePtr *Node
		nodePtr, err = UnmarshalNodePtrJSON(fields["parent_ec"])
		if err != nil {
			return
		}
		if nodePtr != nil && *nodePtr != nil {
			val := (*nodePtr).(EquivalenceClass)
			node.ParentEc = &val
		}
	}

	if fields["eval_cost"] != nil {
		err = json.Unmarshal(fields["eval_cost"], &node.EvalCost)
		if err != nil {
			return
		}
	}

	if fields["norm_selec"] != nil {
		err = json.Unmarshal(fields["norm_selec"], &node.NormSelec)
		if err != nil {
			return
		}
	}

	if fields["outer_selec"] != nil {
		err = json.Unmarshal(fields["outer_selec"], &node.OuterSelec)
		if err != nil {
			return
		}
	}

	if fields["mergeopfamilies"] != nil {
		node.Mergeopfamilies, err = UnmarshalNodeArrayJSON(fields["mergeopfamilies"])
		if err != nil {
			return
		}
	}

	if fields["left_ec"] != nil {
		var nodePtr *Node
		nodePtr, err = UnmarshalNodePtrJSON(fields["left_ec"])
		if err != nil {
			return
		}
		if nodePtr != nil && *nodePtr != nil {
			val := (*nodePtr).(EquivalenceClass)
			node.LeftEc = &val
		}
	}

	if fields["right_ec"] != nil {
		var nodePtr *Node
		nodePtr, err = UnmarshalNodePtrJSON(fields["right_ec"])
		if err != nil {
			return
		}
		if nodePtr != nil && *nodePtr != nil {
			val := (*nodePtr).(EquivalenceClass)
			node.RightEc = &val
		}
	}

	if fields["left_em"] != nil {
		var nodePtr *Node
		nodePtr, err = UnmarshalNodePtrJSON(fields["left_em"])
		if err != nil {
			return
		}
		if nodePtr != nil && *nodePtr != nil {
			val := (*nodePtr).(EquivalenceMember)
			node.LeftEm = &val
		}
	}

	if fields["right_em"] != nil {
		var nodePtr *Node
		nodePtr, err = UnmarshalNodePtrJSON(fields["right_em"])
		if err != nil {
			return
		}
		if nodePtr != nil && *nodePtr != nil {
			val := (*nodePtr).(EquivalenceMember)
			node.RightEm = &val
		}
	}

	if fields["scansel_cache"] != nil {
		node.ScanselCache, err = UnmarshalNodeArrayJSON(fields["scansel_cache"])
		if err != nil {
			return
		}
	}

	if fields["outer_is_left"] != nil {
		err = json.Unmarshal(fields["outer_is_left"], &node.OuterIsLeft)
		if err != nil {
			return
		}
	}

	if fields["hashjoinoperator"] != nil {
		err = json.Unmarshal(fields["hashjoinoperator"], &node.Hashjoinoperator)
		if err != nil {
			return
		}
	}

	if fields["left_bucketsize"] != nil {
		err = json.Unmarshal(fields["left_bucketsize"], &node.LeftBucketsize)
		if err != nil {
			return
		}
	}

	if fields["right_bucketsize"] != nil {
		err = json.Unmarshal(fields["right_bucketsize"], &node.RightBucketsize)
		if err != nil {
			return
		}
	}

	return
}
