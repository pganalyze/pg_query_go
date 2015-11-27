// Auto-generated from postgres/src/include/nodes/relation.h - DO NOT EDIT

package pg_query

import "encoding/json"

/*----------
 * PlannerInfo
 *		Per-query information for planning/optimization
 *
 * This struct is conventionally called "root" in all the planner routines.
 * It holds links to all of the planner's working state, in addition to the
 * original Query.  Note that at present the planner extensively modifies
 * the passed-in Query data structure; someday that should stop.
 *----------
 */
type PlannerInfo struct {
	Parse *Query `json:"parse"` /* the Query being planned */

	Glob *PlannerGlobal `json:"glob"` /* global info for current planner run */

	QueryLevel Index `json:"query_level"` /* 1 at the outermost Query */

	ParentRoot *PlannerInfo `json:"parent_root"` /* NULL at outermost Query */

	PlanParams []Node `json:"plan_params"` /* list of PlannerParamItems, see below */

	/*
	 * simple_rel_array holds pointers to "base rels" and "other rels" (see
	 * comments for RelOptInfo for more info).  It is indexed by rangetable
	 * index (so entry 0 is always wasted).  Entries can be NULL when an RTE
	 * does not correspond to a base relation, such as a join RTE or an
	 * unreferenced view RTE; or if the RelOptInfo hasn't been made yet.
	 */
	SimpleRelArray     *RelOptInfo `json:"simple_rel_array"`      /* All 1-rel RelOptInfos */
	SimpleRelArraySize int         `json:"simple_rel_array_size"` /* allocated size of array */

	/*
	 * simple_rte_array is the same length as simple_rel_array and holds
	 * pointers to the associated rangetable entries.  This lets us avoid
	 * rt_fetch(), which can be a bit slow once large inheritance sets have
	 * been expanded.
	 */
	SimpleRteArray *RangeTblEntry `json:"simple_rte_array"` /* rangetable as an array */

	/*
	 * all_baserels is a Relids set of all base relids (but not "other"
	 * relids) in the query; that is, the Relids identifier of the final join
	 * we need to form.  This is computed in make_one_rel, just before we
	 * start making Paths.
	 */
	AllBaserels []uint32 `json:"all_baserels"`

	/*
	 * nullable_baserels is a Relids set of base relids that are nullable by
	 * some outer join in the jointree; these are rels that are potentially
	 * nullable below the WHERE clause, SELECT targetlist, etc.  This is
	 * computed in deconstruct_jointree.
	 */
	NullableBaserels []uint32 `json:"nullable_baserels"`

	/*
	 * join_rel_list is a list of all join-relation RelOptInfos we have
	 * considered in this planning run.  For small problems we just scan the
	 * list to do lookups, but when there are many join relations we build a
	 * hash table for faster lookups.  The hash table is present and valid
	 * when join_rel_hash is not NULL.  Note that we still maintain the list
	 * even when using the hash table for lookups; this simplifies life for
	 * GEQO.
	 */
	JoinRelList []Node `json:"join_rel_list"` /* list of join-relation RelOptInfos */

	/*
	 * When doing a dynamic-programming-style join search, join_rel_level[k]
	 * is a list of all join-relation RelOptInfos of level k, and
	 * join_cur_level is the current level.  New join-relation RelOptInfos are
	 * automatically added to the join_rel_level[join_cur_level] list.
	 * join_rel_level is NULL if not in use.
	 */
	JoinRelLevel []Node `json:"join_rel_level"` /* lists of join-relation RelOptInfos */
	JoinCurLevel int    `json:"join_cur_level"` /* index of list being extended */

	InitPlans []Node `json:"init_plans"` /* init SubPlans for query */

	CtePlanIds []Node `json:"cte_plan_ids"` /* per-CTE-item list of subplan IDs */

	EqClasses []Node `json:"eq_classes"` /* list of active EquivalenceClasses */

	CanonPathkeys []Node `json:"canon_pathkeys"` /* list of "canonical" PathKeys */

	LeftJoinClauses []Node `json:"left_join_clauses"` /* list of RestrictInfos for
	 * mergejoinable outer join clauses
	 * w/nonnullable var on left */

	RightJoinClauses []Node `json:"right_join_clauses"` /* list of RestrictInfos for
	 * mergejoinable outer join clauses
	 * w/nonnullable var on right */

	FullJoinClauses []Node `json:"full_join_clauses"` /* list of RestrictInfos for
	 * mergejoinable full join clauses */

	JoinInfoList []Node `json:"join_info_list"` /* list of SpecialJoinInfos */

	LateralInfoList []Node `json:"lateral_info_list"` /* list of LateralJoinInfos */

	AppendRelList []Node `json:"append_rel_list"` /* list of AppendRelInfos */

	RowMarks []Node `json:"rowMarks"` /* list of PlanRowMarks */

	PlaceholderList []Node `json:"placeholder_list"` /* list of PlaceHolderInfos */

	QueryPathkeys []Node `json:"query_pathkeys"` /* desired pathkeys for query_planner(), and
	 * actual pathkeys after planning */

	GroupPathkeys    []Node `json:"group_pathkeys"`    /* groupClause pathkeys, if any */
	WindowPathkeys   []Node `json:"window_pathkeys"`   /* pathkeys of bottom window, if any */
	DistinctPathkeys []Node `json:"distinct_pathkeys"` /* distinctClause pathkeys, if any */
	SortPathkeys     []Node `json:"sort_pathkeys"`     /* sortClause pathkeys, if any */

	MinmaxAggs []Node `json:"minmax_aggs"` /* List of MinMaxAggInfos */

	InitialRels []Node `json:"initial_rels"` /* RelOptInfos we are now trying to join */

	TotalTablePages float64 `json:"total_table_pages"` /* # of pages in all tables of query */

	TupleFraction float64 `json:"tuple_fraction"` /* tuple_fraction passed to query_planner */
	LimitTuples   float64 `json:"limit_tuples"`   /* limit_tuples passed to query_planner */

	HasInheritedTarget bool `json:"hasInheritedTarget"` /* true if parse->resultRelation is an
	 * inheritance child rel */

	HasJoinRtes            bool `json:"hasJoinRTEs"`            /* true if any RTEs are RTE_JOIN kind */
	HasLateralRtes         bool `json:"hasLateralRTEs"`         /* true if any RTEs are marked LATERAL */
	HasHavingQual          bool `json:"hasHavingQual"`          /* true if havingQual was non-null */
	HasPseudoConstantQuals bool `json:"hasPseudoConstantQuals"` /* true if any RestrictInfo has
	 * pseudoconstant = true */

	HasRecursion bool `json:"hasRecursion"` /* true if planning a recursive WITH item */

	/* These fields are used only when hasRecursion is true: */
	WtParamId        int   `json:"wt_param_id"`        /* PARAM_EXEC ID for the work table */
	NonRecursivePlan *Plan `json:"non_recursive_plan"` /* plan for non-recursive term */

	/* These fields are workspace for createplan.c */
	CurOuterRels   []uint32 `json:"curOuterRels"`   /* outer rels above current node */
	CurOuterParams []Node   `json:"curOuterParams"` /* not-yet-assigned NestLoopParams */

	/* optional private data for join_search_hook, e.g., GEQO */
	JoinSearchPrivate interface{} `json:"join_search_private"`
}

func (node PlannerInfo) MarshalJSON() ([]byte, error) {
	type PlannerInfoMarshalAlias PlannerInfo
	return json.Marshal(map[string]interface{}{
		"PLANNERINFO": (*PlannerInfoMarshalAlias)(&node),
	})
}

func (node *PlannerInfo) UnmarshalJSON(input []byte) (err error) {
	var fields map[string]json.RawMessage

	err = json.Unmarshal(input, &fields)
	if err != nil {
		return
	}

	if fields["parse"] != nil {
		var nodePtr *Node
		nodePtr, err = UnmarshalNodePtrJSON(fields["parse"])
		if err != nil {
			return
		}
		if nodePtr != nil && *nodePtr != nil {
			val := (*nodePtr).(Query)
			node.Parse = &val
		}
	}

	if fields["glob"] != nil {
		var nodePtr *Node
		nodePtr, err = UnmarshalNodePtrJSON(fields["glob"])
		if err != nil {
			return
		}
		if nodePtr != nil && *nodePtr != nil {
			val := (*nodePtr).(PlannerGlobal)
			node.Glob = &val
		}
	}

	if fields["query_level"] != nil {
		err = json.Unmarshal(fields["query_level"], &node.QueryLevel)
		if err != nil {
			return
		}
	}

	if fields["parent_root"] != nil {
		var nodePtr *Node
		nodePtr, err = UnmarshalNodePtrJSON(fields["parent_root"])
		if err != nil {
			return
		}
		if nodePtr != nil && *nodePtr != nil {
			val := (*nodePtr).(PlannerInfo)
			node.ParentRoot = &val
		}
	}

	if fields["plan_params"] != nil {
		node.PlanParams, err = UnmarshalNodeArrayJSON(fields["plan_params"])
		if err != nil {
			return
		}
	}

	if fields["simple_rel_array"] != nil {
		var nodePtr *Node
		nodePtr, err = UnmarshalNodePtrJSON(fields["simple_rel_array"])
		if err != nil {
			return
		}
		if nodePtr != nil && *nodePtr != nil {
			val := (*nodePtr).(RelOptInfo)
			node.SimpleRelArray = &val
		}
	}

	if fields["simple_rel_array_size"] != nil {
		err = json.Unmarshal(fields["simple_rel_array_size"], &node.SimpleRelArraySize)
		if err != nil {
			return
		}
	}

	if fields["simple_rte_array"] != nil {
		var nodePtr *Node
		nodePtr, err = UnmarshalNodePtrJSON(fields["simple_rte_array"])
		if err != nil {
			return
		}
		if nodePtr != nil && *nodePtr != nil {
			val := (*nodePtr).(RangeTblEntry)
			node.SimpleRteArray = &val
		}
	}

	if fields["all_baserels"] != nil {
		err = json.Unmarshal(fields["all_baserels"], &node.AllBaserels)
		if err != nil {
			return
		}
	}

	if fields["nullable_baserels"] != nil {
		err = json.Unmarshal(fields["nullable_baserels"], &node.NullableBaserels)
		if err != nil {
			return
		}
	}

	if fields["join_rel_list"] != nil {
		node.JoinRelList, err = UnmarshalNodeArrayJSON(fields["join_rel_list"])
		if err != nil {
			return
		}
	}

	if fields["join_rel_level"] != nil {
		node.JoinRelLevel, err = UnmarshalNodeArrayJSON(fields["join_rel_level"])
		if err != nil {
			return
		}
	}

	if fields["join_cur_level"] != nil {
		err = json.Unmarshal(fields["join_cur_level"], &node.JoinCurLevel)
		if err != nil {
			return
		}
	}

	if fields["init_plans"] != nil {
		node.InitPlans, err = UnmarshalNodeArrayJSON(fields["init_plans"])
		if err != nil {
			return
		}
	}

	if fields["cte_plan_ids"] != nil {
		node.CtePlanIds, err = UnmarshalNodeArrayJSON(fields["cte_plan_ids"])
		if err != nil {
			return
		}
	}

	if fields["eq_classes"] != nil {
		node.EqClasses, err = UnmarshalNodeArrayJSON(fields["eq_classes"])
		if err != nil {
			return
		}
	}

	if fields["canon_pathkeys"] != nil {
		node.CanonPathkeys, err = UnmarshalNodeArrayJSON(fields["canon_pathkeys"])
		if err != nil {
			return
		}
	}

	if fields["left_join_clauses"] != nil {
		node.LeftJoinClauses, err = UnmarshalNodeArrayJSON(fields["left_join_clauses"])
		if err != nil {
			return
		}
	}

	if fields["right_join_clauses"] != nil {
		node.RightJoinClauses, err = UnmarshalNodeArrayJSON(fields["right_join_clauses"])
		if err != nil {
			return
		}
	}

	if fields["full_join_clauses"] != nil {
		node.FullJoinClauses, err = UnmarshalNodeArrayJSON(fields["full_join_clauses"])
		if err != nil {
			return
		}
	}

	if fields["join_info_list"] != nil {
		node.JoinInfoList, err = UnmarshalNodeArrayJSON(fields["join_info_list"])
		if err != nil {
			return
		}
	}

	if fields["lateral_info_list"] != nil {
		node.LateralInfoList, err = UnmarshalNodeArrayJSON(fields["lateral_info_list"])
		if err != nil {
			return
		}
	}

	if fields["append_rel_list"] != nil {
		node.AppendRelList, err = UnmarshalNodeArrayJSON(fields["append_rel_list"])
		if err != nil {
			return
		}
	}

	if fields["rowMarks"] != nil {
		node.RowMarks, err = UnmarshalNodeArrayJSON(fields["rowMarks"])
		if err != nil {
			return
		}
	}

	if fields["placeholder_list"] != nil {
		node.PlaceholderList, err = UnmarshalNodeArrayJSON(fields["placeholder_list"])
		if err != nil {
			return
		}
	}

	if fields["query_pathkeys"] != nil {
		node.QueryPathkeys, err = UnmarshalNodeArrayJSON(fields["query_pathkeys"])
		if err != nil {
			return
		}
	}

	if fields["group_pathkeys"] != nil {
		node.GroupPathkeys, err = UnmarshalNodeArrayJSON(fields["group_pathkeys"])
		if err != nil {
			return
		}
	}

	if fields["window_pathkeys"] != nil {
		node.WindowPathkeys, err = UnmarshalNodeArrayJSON(fields["window_pathkeys"])
		if err != nil {
			return
		}
	}

	if fields["distinct_pathkeys"] != nil {
		node.DistinctPathkeys, err = UnmarshalNodeArrayJSON(fields["distinct_pathkeys"])
		if err != nil {
			return
		}
	}

	if fields["sort_pathkeys"] != nil {
		node.SortPathkeys, err = UnmarshalNodeArrayJSON(fields["sort_pathkeys"])
		if err != nil {
			return
		}
	}

	if fields["minmax_aggs"] != nil {
		node.MinmaxAggs, err = UnmarshalNodeArrayJSON(fields["minmax_aggs"])
		if err != nil {
			return
		}
	}

	if fields["initial_rels"] != nil {
		node.InitialRels, err = UnmarshalNodeArrayJSON(fields["initial_rels"])
		if err != nil {
			return
		}
	}

	if fields["total_table_pages"] != nil {
		err = json.Unmarshal(fields["total_table_pages"], &node.TotalTablePages)
		if err != nil {
			return
		}
	}

	if fields["tuple_fraction"] != nil {
		err = json.Unmarshal(fields["tuple_fraction"], &node.TupleFraction)
		if err != nil {
			return
		}
	}

	if fields["limit_tuples"] != nil {
		err = json.Unmarshal(fields["limit_tuples"], &node.LimitTuples)
		if err != nil {
			return
		}
	}

	if fields["hasInheritedTarget"] != nil {
		err = json.Unmarshal(fields["hasInheritedTarget"], &node.HasInheritedTarget)
		if err != nil {
			return
		}
	}

	if fields["hasJoinRTEs"] != nil {
		err = json.Unmarshal(fields["hasJoinRTEs"], &node.HasJoinRtes)
		if err != nil {
			return
		}
	}

	if fields["hasLateralRTEs"] != nil {
		err = json.Unmarshal(fields["hasLateralRTEs"], &node.HasLateralRtes)
		if err != nil {
			return
		}
	}

	if fields["hasHavingQual"] != nil {
		err = json.Unmarshal(fields["hasHavingQual"], &node.HasHavingQual)
		if err != nil {
			return
		}
	}

	if fields["hasPseudoConstantQuals"] != nil {
		err = json.Unmarshal(fields["hasPseudoConstantQuals"], &node.HasPseudoConstantQuals)
		if err != nil {
			return
		}
	}

	if fields["hasRecursion"] != nil {
		err = json.Unmarshal(fields["hasRecursion"], &node.HasRecursion)
		if err != nil {
			return
		}
	}

	if fields["wt_param_id"] != nil {
		err = json.Unmarshal(fields["wt_param_id"], &node.WtParamId)
		if err != nil {
			return
		}
	}

	if fields["non_recursive_plan"] != nil {
		var nodePtr *Node
		nodePtr, err = UnmarshalNodePtrJSON(fields["non_recursive_plan"])
		if err != nil {
			return
		}
		if nodePtr != nil && *nodePtr != nil {
			val := (*nodePtr).(Plan)
			node.NonRecursivePlan = &val
		}
	}

	if fields["curOuterRels"] != nil {
		err = json.Unmarshal(fields["curOuterRels"], &node.CurOuterRels)
		if err != nil {
			return
		}
	}

	if fields["curOuterParams"] != nil {
		node.CurOuterParams, err = UnmarshalNodeArrayJSON(fields["curOuterParams"])
		if err != nil {
			return
		}
	}

	if fields["join_search_private"] != nil {
		err = json.Unmarshal(fields["join_search_private"], &node.JoinSearchPrivate)
		if err != nil {
			return
		}
	}

	return
}
