// Auto-generated from postgres/src/include/nodes/relation.h - DO NOT EDIT

package pg_query

import "encoding/json"

/*
 * Restriction clause info.
 *
 * We create one of these for each AND sub-clause of a restriction condition
 * (WHERE or JOIN/ON clause).  Since the restriction clauses are logically
 * ANDed, we can use any one of them or any subset of them to filter out
 * tuples, without having to evaluate the rest.  The RestrictInfo node itself
 * stores data used by the optimizer while choosing the best query plan.
 *
 * If a restriction clause references a single base relation, it will appear
 * in the baserestrictinfo list of the RelOptInfo for that base rel.
 *
 * If a restriction clause references more than one base rel, it will
 * appear in the joininfo list of every RelOptInfo that describes a strict
 * subset of the base rels mentioned in the clause.  The joininfo lists are
 * used to drive join tree building by selecting plausible join candidates.
 * The clause cannot actually be applied until we have built a join rel
 * containing all the base rels it references, however.
 *
 * When we construct a join rel that includes all the base rels referenced
 * in a multi-relation restriction clause, we place that clause into the
 * joinrestrictinfo lists of paths for the join rel, if neither left nor
 * right sub-path includes all base rels referenced in the clause.  The clause
 * will be applied at that join level, and will not propagate any further up
 * the join tree.  (Note: the "predicate migration" code was once intended to
 * push restriction clauses up and down the plan tree based on evaluation
 * costs, but it's dead code and is unlikely to be resurrected in the
 * foreseeable future.)
 *
 * Note that in the presence of more than two rels, a multi-rel restriction
 * might reach different heights in the join tree depending on the join
 * sequence we use.  So, these clauses cannot be associated directly with
 * the join RelOptInfo, but must be kept track of on a per-join-path basis.
 *
 * RestrictInfos that represent equivalence conditions (i.e., mergejoinable
 * equalities that are not outerjoin-delayed) are handled a bit differently.
 * Initially we attach them to the EquivalenceClasses that are derived from
 * them.  When we construct a scan or join path, we look through all the
 * EquivalenceClasses and generate derived RestrictInfos representing the
 * minimal set of conditions that need to be checked for this particular scan
 * or join to enforce that all members of each EquivalenceClass are in fact
 * equal in all rows emitted by the scan or join.
 *
 * When dealing with outer joins we have to be very careful about pushing qual
 * clauses up and down the tree.  An outer join's own JOIN/ON conditions must
 * be evaluated exactly at that join node, unless they are "degenerate"
 * conditions that reference only Vars from the nullable side of the join.
 * Quals appearing in WHERE or in a JOIN above the outer join cannot be pushed
 * down below the outer join, if they reference any nullable Vars.
 * RestrictInfo nodes contain a flag to indicate whether a qual has been
 * pushed down to a lower level than its original syntactic placement in the
 * join tree would suggest.  If an outer join prevents us from pushing a qual
 * down to its "natural" semantic level (the level associated with just the
 * base rels used in the qual) then we mark the qual with a "required_relids"
 * value including more than just the base rels it actually uses.  By
 * pretending that the qual references all the rels required to form the outer
 * join, we prevent it from being evaluated below the outer join's joinrel.
 * When we do form the outer join's joinrel, we still need to distinguish
 * those quals that are actually in that join's JOIN/ON condition from those
 * that appeared elsewhere in the tree and were pushed down to the join rel
 * because they used no other rels.  That's what the is_pushed_down flag is
 * for; it tells us that a qual is not an OUTER JOIN qual for the set of base
 * rels listed in required_relids.  A clause that originally came from WHERE
 * or an INNER JOIN condition will *always* have its is_pushed_down flag set.
 * It's possible for an OUTER JOIN clause to be marked is_pushed_down too,
 * if we decide that it can be pushed down into the nullable side of the join.
 * In that case it acts as a plain filter qual for wherever it gets evaluated.
 * (In short, is_pushed_down is only false for non-degenerate outer join
 * conditions.  Possibly we should rename it to reflect that meaning?)
 *
 * RestrictInfo nodes also contain an outerjoin_delayed flag, which is true
 * if the clause's applicability must be delayed due to any outer joins
 * appearing below it (ie, it has to be postponed to some join level higher
 * than the set of relations it actually references).
 *
 * There is also an outer_relids field, which is NULL except for outer join
 * clauses; for those, it is the set of relids on the outer side of the
 * clause's outer join.  (These are rels that the clause cannot be applied to
 * in parameterized scans, since pushing it into the join's outer side would
 * lead to wrong answers.)
 *
 * There is also a nullable_relids field, which is the set of rels the clause
 * references that can be forced null by some outer join below the clause.
 *
 * outerjoin_delayed = true is subtly different from nullable_relids != NULL:
 * a clause might reference some nullable rels and yet not be
 * outerjoin_delayed because it also references all the other rels of the
 * outer join(s). A clause that is not outerjoin_delayed can be enforced
 * anywhere it is computable.
 *
 * In general, the referenced clause might be arbitrarily complex.  The
 * kinds of clauses we can handle as indexscan quals, mergejoin clauses,
 * or hashjoin clauses are limited (e.g., no volatile functions).  The code
 * for each kind of path is responsible for identifying the restrict clauses
 * it can use and ignoring the rest.  Clauses not implemented by an indexscan,
 * mergejoin, or hashjoin will be placed in the plan qual or joinqual field
 * of the finished Plan node, where they will be enforced by general-purpose
 * qual-expression-evaluation code.  (But we are still entitled to count
 * their selectivity when estimating the result tuple count, if we
 * can guess what it is...)
 *
 * When the referenced clause is an OR clause, we generate a modified copy
 * in which additional RestrictInfo nodes are inserted below the top-level
 * OR/AND structure.  This is a convenience for OR indexscan processing:
 * indexquals taken from either the top level or an OR subclause will have
 * associated RestrictInfo nodes.
 *
 * The can_join flag is set true if the clause looks potentially useful as
 * a merge or hash join clause, that is if it is a binary opclause with
 * nonoverlapping sets of relids referenced in the left and right sides.
 * (Whether the operator is actually merge or hash joinable isn't checked,
 * however.)
 *
 * The pseudoconstant flag is set true if the clause contains no Vars of
 * the current query level and no volatile functions.  Such a clause can be
 * pulled out and used as a one-time qual in a gating Result node.  We keep
 * pseudoconstant clauses in the same lists as other RestrictInfos so that
 * the regular clause-pushing machinery can assign them to the correct join
 * level, but they need to be treated specially for cost and selectivity
 * estimates.  Note that a pseudoconstant clause can never be an indexqual
 * or merge or hash join clause, so it's of no interest to large parts of
 * the planner.
 *
 * When join clauses are generated from EquivalenceClasses, there may be
 * several equally valid ways to enforce join equivalence, of which we need
 * apply only one.  We mark clauses of this kind by setting parent_ec to
 * point to the generating EquivalenceClass.  Multiple clauses with the same
 * parent_ec in the same join are redundant.
 */
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
