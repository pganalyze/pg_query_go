// Auto-generated from postgres/src/include/nodes/relation.h - DO NOT EDIT

package pg_query

import "encoding/json"

/*----------
 * RelOptInfo
 *		Per-relation information for planning/optimization
 *
 * For planning purposes, a "base rel" is either a plain relation (a table)
 * or the output of a sub-SELECT or function that appears in the range table.
 * In either case it is uniquely identified by an RT index.  A "joinrel"
 * is the joining of two or more base rels.  A joinrel is identified by
 * the set of RT indexes for its component baserels.  We create RelOptInfo
 * nodes for each baserel and joinrel, and store them in the PlannerInfo's
 * simple_rel_array and join_rel_list respectively.
 *
 * Note that there is only one joinrel for any given set of component
 * baserels, no matter what order we assemble them in; so an unordered
 * set is the right datatype to identify it with.
 *
 * We also have "other rels", which are like base rels in that they refer to
 * single RT indexes; but they are not part of the join tree, and are given
 * a different RelOptKind to identify them.  Lastly, there is a RelOptKind
 * for "dead" relations, which are base rels that we have proven we don't
 * need to join after all.
 *
 * Currently the only kind of otherrels are those made for member relations
 * of an "append relation", that is an inheritance set or UNION ALL subquery.
 * An append relation has a parent RTE that is a base rel, which represents
 * the entire append relation.  The member RTEs are otherrels.  The parent
 * is present in the query join tree but the members are not.  The member
 * RTEs and otherrels are used to plan the scans of the individual tables or
 * subqueries of the append set; then the parent baserel is given Append
 * and/or MergeAppend paths comprising the best paths for the individual
 * member rels.  (See comments for AppendRelInfo for more information.)
 *
 * At one time we also made otherrels to represent join RTEs, for use in
 * handling join alias Vars.  Currently this is not needed because all join
 * alias Vars are expanded to non-aliased form during preprocess_expression.
 *
 * Parts of this data structure are specific to various scan and join
 * mechanisms.  It didn't seem worth creating new node types for them.
 *
 *		relids - Set of base-relation identifiers; it is a base relation
 *				if there is just one, a join relation if more than one
 *		rows - estimated number of tuples in the relation after restriction
 *			   clauses have been applied (ie, output rows of a plan for it)
 *		width - avg. number of bytes per tuple in the relation after the
 *				appropriate projections have been done (ie, output width)
 *		consider_startup - true if there is any value in keeping plain paths for
 *						   this rel on the basis of having cheap startup cost
 *		consider_param_startup - the same for parameterized paths
 *		reltargetlist - List of Var and PlaceHolderVar nodes for the values
 *						we need to output from this relation.
 *						List is in no particular order, but all rels of an
 *						appendrel set must use corresponding orders.
 *						NOTE: in an appendrel child relation, may contain
 *						arbitrary expressions pulled up from a subquery!
 *		pathlist - List of Path nodes, one for each potentially useful
 *				   method of generating the relation
 *		ppilist - ParamPathInfo nodes for parameterized Paths, if any
 *		cheapest_startup_path - the pathlist member with lowest startup cost
 *			(regardless of ordering) among the unparameterized paths;
 *			or NULL if there is no unparameterized path
 *		cheapest_total_path - the pathlist member with lowest total cost
 *			(regardless of ordering) among the unparameterized paths;
 *			or if there is no unparameterized path, the path with lowest
 *			total cost among the paths with minimum parameterization
 *		cheapest_unique_path - for caching cheapest path to produce unique
 *			(no duplicates) output from relation; NULL if not yet requested
 *		cheapest_parameterized_paths - best paths for their parameterizations;
 *			always includes cheapest_total_path, even if that's unparameterized
 *
 * If the relation is a base relation it will have these fields set:
 *
 *		relid - RTE index (this is redundant with the relids field, but
 *				is provided for convenience of access)
 *		rtekind - distinguishes plain relation, subquery, or function RTE
 *		min_attr, max_attr - range of valid AttrNumbers for rel
 *		attr_needed - array of bitmapsets indicating the highest joinrel
 *				in which each attribute is needed; if bit 0 is set then
 *				the attribute is needed as part of final targetlist
 *		attr_widths - cache space for per-attribute width estimates;
 *					  zero means not computed yet
 *		lateral_vars - lateral cross-references of rel, if any (list of
 *					   Vars and PlaceHolderVars)
 *		lateral_relids - required outer rels for LATERAL, as a Relids set
 *						 (for child rels this can be more than lateral_vars)
 *		lateral_referencers - relids of rels that reference this one laterally
 *		indexlist - list of IndexOptInfo nodes for relation's indexes
 *					(always NIL if it's not a table)
 *		pages - number of disk pages in relation (zero if not a table)
 *		tuples - number of tuples in relation (not considering restrictions)
 *		allvisfrac - fraction of disk pages that are marked all-visible
 *		subplan - plan for subquery (NULL if it's not a subquery)
 *		subroot - PlannerInfo for subquery (NULL if it's not a subquery)
 *		subplan_params - list of PlannerParamItems to be passed to subquery
 *		fdwroutine - function hooks for FDW, if foreign table (else NULL)
 *		fdw_private - private state for FDW, if foreign table (else NULL)
 *
 *		Note: for a subquery, tuples, subplan, subroot are not set immediately
 *		upon creation of the RelOptInfo object; they are filled in when
 *		set_subquery_pathlist processes the object.  Likewise, fdwroutine
 *		and fdw_private are filled during initial path creation.
 *
 *		For otherrels that are appendrel members, these fields are filled
 *		in just as for a baserel.
 *
 * The presence of the remaining fields depends on the restrictions
 * and joins that the relation participates in:
 *
 *		baserestrictinfo - List of RestrictInfo nodes, containing info about
 *					each non-join qualification clause in which this relation
 *					participates (only used for base rels)
 *		baserestrictcost - Estimated cost of evaluating the baserestrictinfo
 *					clauses at a single tuple (only used for base rels)
 *		joininfo  - List of RestrictInfo nodes, containing info about each
 *					join clause in which this relation participates (but
 *					note this excludes clauses that might be derivable from
 *					EquivalenceClasses)
 *		has_eclass_joins - flag that EquivalenceClass joins are possible
 *
 * Note: Keeping a restrictinfo list in the RelOptInfo is useful only for
 * base rels, because for a join rel the set of clauses that are treated as
 * restrict clauses varies depending on which sub-relations we choose to join.
 * (For example, in a 3-base-rel join, a clause relating rels 1 and 2 must be
 * treated as a restrictclause if we join {1} and {2 3} to make {1 2 3}; but
 * if we join {1 2} and {3} then that clause will be a restrictclause in {1 2}
 * and should not be processed again at the level of {1 2 3}.)	Therefore,
 * the restrictinfo list in the join case appears in individual JoinPaths
 * (field joinrestrictinfo), not in the parent relation.  But it's OK for
 * the RelOptInfo to store the joininfo list, because that is the same
 * for a given rel no matter how we form it.
 *
 * We store baserestrictcost in the RelOptInfo (for base relations) because
 * we know we will need it at least once (to price the sequential scan)
 * and may need it multiple times to price index scans.
 *----------
 */
type RelOptInfo struct {
	Reloptkind RelOptKind `json:"reloptkind"`

	/* all relations included in this RelOptInfo */
	Relids []uint32 `json:"relids"` /* set of base relids (rangetable indexes) */

	/* size estimates generated by planner */
	Rows  float64 `json:"rows"`  /* estimated number of result tuples */
	Width int     `json:"width"` /* estimated avg width of result tuples */

	/* per-relation planner control flags */
	ConsiderStartup      bool `json:"consider_startup"`       /* keep cheap-startup-cost paths? */
	ConsiderParamStartup bool `json:"consider_param_startup"` /* ditto, for parameterized paths? */

	/* materialization information */
	Reltargetlist              []Node `json:"reltargetlist"` /* Vars to be output by scan of relation */
	Pathlist                   []Node `json:"pathlist"`      /* Path structures */
	Ppilist                    []Node `json:"ppilist"`       /* ParamPathInfos used in pathlist */
	CheapestStartupPath        *Path  `json:"cheapest_startup_path"`
	CheapestTotalPath          *Path  `json:"cheapest_total_path"`
	CheapestUniquePath         *Path  `json:"cheapest_unique_path"`
	CheapestParameterizedPaths []Node `json:"cheapest_parameterized_paths"`

	/* information about a base rel (not set for join rels!) */
	Relid              Index       `json:"relid"`
	Reltablespace      Oid         `json:"reltablespace"`       /* containing tablespace */
	Rtekind            RTEKind     `json:"rtekind"`             /* RELATION, SUBQUERY, or FUNCTION */
	MinAttr            AttrNumber  `json:"min_attr"`            /* smallest attrno of rel (often <0) */
	MaxAttr            AttrNumber  `json:"max_attr"`            /* largest attrno of rel */
	AttrNeeded         []Relids    `json:"attr_needed"`         /* array indexed [min_attr .. max_attr] */
	AttrWidths         []int32     `json:"attr_widths"`         /* array indexed [min_attr .. max_attr] */
	LateralVars        []Node      `json:"lateral_vars"`        /* LATERAL Vars and PHVs referenced by rel */
	LateralRelids      []uint32    `json:"lateral_relids"`      /* minimum parameterization of rel */
	LateralReferencers []uint32    `json:"lateral_referencers"` /* rels that reference me laterally */
	Indexlist          []Node      `json:"indexlist"`           /* list of IndexOptInfo */
	Pages              BlockNumber `json:"pages"`               /* size estimates derived from pg_class */
	Tuples             float64     `json:"tuples"`
	Allvisfrac         float64     `json:"allvisfrac"`

	/* use "struct Plan" to avoid including plannodes.h here */
	Subplan       *Plan        `json:"subplan"`        /* if subquery */
	Subroot       *PlannerInfo `json:"subroot"`        /* if subquery */
	SubplanParams []Node       `json:"subplan_params"` /* if subquery */

	/* use "struct FdwRoutine" to avoid including fdwapi.h here */
	FdwPrivate interface{} `json:"fdw_private"` /* if foreign table */

	/* used by various scans and joins: */
	Baserestrictinfo []Node `json:"baserestrictinfo"` /* RestrictInfo structures (if base
	 * rel) */

	Baserestrictcost QualCost `json:"baserestrictcost"` /* cost of evaluating the above */
	Joininfo         []Node   `json:"joininfo"`         /* RestrictInfo structures for join clauses
	 * involving this rel */

	HasEclassJoins bool `json:"has_eclass_joins"` /* T means joininfo is incomplete */
}

func (node RelOptInfo) MarshalJSON() ([]byte, error) {
	type RelOptInfoMarshalAlias RelOptInfo
	return json.Marshal(map[string]interface{}{
		"RELOPTINFO": (*RelOptInfoMarshalAlias)(&node),
	})
}

func (node *RelOptInfo) UnmarshalJSON(input []byte) (err error) {
	var fields map[string]json.RawMessage

	err = json.Unmarshal(input, &fields)
	if err != nil {
		return
	}

	if fields["reloptkind"] != nil {
		err = json.Unmarshal(fields["reloptkind"], &node.Reloptkind)
		if err != nil {
			return
		}
	}

	if fields["relids"] != nil {
		err = json.Unmarshal(fields["relids"], &node.Relids)
		if err != nil {
			return
		}
	}

	if fields["rows"] != nil {
		err = json.Unmarshal(fields["rows"], &node.Rows)
		if err != nil {
			return
		}
	}

	if fields["width"] != nil {
		err = json.Unmarshal(fields["width"], &node.Width)
		if err != nil {
			return
		}
	}

	if fields["consider_startup"] != nil {
		err = json.Unmarshal(fields["consider_startup"], &node.ConsiderStartup)
		if err != nil {
			return
		}
	}

	if fields["consider_param_startup"] != nil {
		err = json.Unmarshal(fields["consider_param_startup"], &node.ConsiderParamStartup)
		if err != nil {
			return
		}
	}

	if fields["reltargetlist"] != nil {
		node.Reltargetlist, err = UnmarshalNodeArrayJSON(fields["reltargetlist"])
		if err != nil {
			return
		}
	}

	if fields["pathlist"] != nil {
		node.Pathlist, err = UnmarshalNodeArrayJSON(fields["pathlist"])
		if err != nil {
			return
		}
	}

	if fields["ppilist"] != nil {
		node.Ppilist, err = UnmarshalNodeArrayJSON(fields["ppilist"])
		if err != nil {
			return
		}
	}

	if fields["cheapest_startup_path"] != nil {
		var nodePtr *Node
		nodePtr, err = UnmarshalNodePtrJSON(fields["cheapest_startup_path"])
		if err != nil {
			return
		}
		if nodePtr != nil && *nodePtr != nil {
			val := (*nodePtr).(Path)
			node.CheapestStartupPath = &val
		}
	}

	if fields["cheapest_total_path"] != nil {
		var nodePtr *Node
		nodePtr, err = UnmarshalNodePtrJSON(fields["cheapest_total_path"])
		if err != nil {
			return
		}
		if nodePtr != nil && *nodePtr != nil {
			val := (*nodePtr).(Path)
			node.CheapestTotalPath = &val
		}
	}

	if fields["cheapest_unique_path"] != nil {
		var nodePtr *Node
		nodePtr, err = UnmarshalNodePtrJSON(fields["cheapest_unique_path"])
		if err != nil {
			return
		}
		if nodePtr != nil && *nodePtr != nil {
			val := (*nodePtr).(Path)
			node.CheapestUniquePath = &val
		}
	}

	if fields["cheapest_parameterized_paths"] != nil {
		node.CheapestParameterizedPaths, err = UnmarshalNodeArrayJSON(fields["cheapest_parameterized_paths"])
		if err != nil {
			return
		}
	}

	if fields["relid"] != nil {
		err = json.Unmarshal(fields["relid"], &node.Relid)
		if err != nil {
			return
		}
	}

	if fields["reltablespace"] != nil {
		err = json.Unmarshal(fields["reltablespace"], &node.Reltablespace)
		if err != nil {
			return
		}
	}

	if fields["rtekind"] != nil {
		err = json.Unmarshal(fields["rtekind"], &node.Rtekind)
		if err != nil {
			return
		}
	}

	if fields["min_attr"] != nil {
		err = json.Unmarshal(fields["min_attr"], &node.MinAttr)
		if err != nil {
			return
		}
	}

	if fields["max_attr"] != nil {
		err = json.Unmarshal(fields["max_attr"], &node.MaxAttr)
		if err != nil {
			return
		}
	}

	if fields["attr_needed"] != nil {
		err = json.Unmarshal(fields["attr_needed"], &node.AttrNeeded)
		if err != nil {
			return
		}
	}

	if fields["attr_widths"] != nil {
		err = json.Unmarshal(fields["attr_widths"], &node.AttrWidths)
		if err != nil {
			return
		}
	}

	if fields["lateral_vars"] != nil {
		node.LateralVars, err = UnmarshalNodeArrayJSON(fields["lateral_vars"])
		if err != nil {
			return
		}
	}

	if fields["lateral_relids"] != nil {
		err = json.Unmarshal(fields["lateral_relids"], &node.LateralRelids)
		if err != nil {
			return
		}
	}

	if fields["lateral_referencers"] != nil {
		err = json.Unmarshal(fields["lateral_referencers"], &node.LateralReferencers)
		if err != nil {
			return
		}
	}

	if fields["indexlist"] != nil {
		node.Indexlist, err = UnmarshalNodeArrayJSON(fields["indexlist"])
		if err != nil {
			return
		}
	}

	if fields["pages"] != nil {
		err = json.Unmarshal(fields["pages"], &node.Pages)
		if err != nil {
			return
		}
	}

	if fields["tuples"] != nil {
		err = json.Unmarshal(fields["tuples"], &node.Tuples)
		if err != nil {
			return
		}
	}

	if fields["allvisfrac"] != nil {
		err = json.Unmarshal(fields["allvisfrac"], &node.Allvisfrac)
		if err != nil {
			return
		}
	}

	if fields["subplan"] != nil {
		var nodePtr *Node
		nodePtr, err = UnmarshalNodePtrJSON(fields["subplan"])
		if err != nil {
			return
		}
		if nodePtr != nil && *nodePtr != nil {
			val := (*nodePtr).(Plan)
			node.Subplan = &val
		}
	}

	if fields["subroot"] != nil {
		var nodePtr *Node
		nodePtr, err = UnmarshalNodePtrJSON(fields["subroot"])
		if err != nil {
			return
		}
		if nodePtr != nil && *nodePtr != nil {
			val := (*nodePtr).(PlannerInfo)
			node.Subroot = &val
		}
	}

	if fields["subplan_params"] != nil {
		node.SubplanParams, err = UnmarshalNodeArrayJSON(fields["subplan_params"])
		if err != nil {
			return
		}
	}

	if fields["fdw_private"] != nil {
		err = json.Unmarshal(fields["fdw_private"], &node.FdwPrivate)
		if err != nil {
			return
		}
	}

	if fields["baserestrictinfo"] != nil {
		node.Baserestrictinfo, err = UnmarshalNodeArrayJSON(fields["baserestrictinfo"])
		if err != nil {
			return
		}
	}

	if fields["baserestrictcost"] != nil {
		err = json.Unmarshal(fields["baserestrictcost"], &node.Baserestrictcost)
		if err != nil {
			return
		}
	}

	if fields["joininfo"] != nil {
		node.Joininfo, err = UnmarshalNodeArrayJSON(fields["joininfo"])
		if err != nil {
			return
		}
	}

	if fields["has_eclass_joins"] != nil {
		err = json.Unmarshal(fields["has_eclass_joins"], &node.HasEclassJoins)
		if err != nil {
			return
		}
	}

	return
}
