// Auto-generated from postgres/src/include/nodes/plannodes.h - DO NOT EDIT

package pg_query

import "encoding/json"

type Plan struct {
	/*
	 * estimated execution costs for plan (see costsize.c for more info)
	 */
	StartupCost Cost `json:"startup_cost"` /* cost expended before fetching any tuples */
	TotalCost   Cost `json:"total_cost"`   /* total cost (assuming all tuples fetched) */

	/*
	 * planner's estimate of result size of this plan step
	 */
	PlanRows  float64 `json:"plan_rows"`  /* number of rows plan is expected to emit */
	PlanWidth int     `json:"plan_width"` /* average row width in bytes */

	/*
	 * Common structural data for all Plan types.
	 */
	Targetlist []Node `json:"targetlist"` /* target list to be computed at this node */
	Qual       []Node `json:"qual"`       /* implicitly-ANDed qual conditions */
	Lefttree   *Plan  `json:"lefttree"`   /* input plan tree(s) */
	Righttree  *Plan  `json:"righttree"`
	InitPlan   []Node `json:"initPlan"` /* Init Plan nodes (un-correlated expr
	 * subselects) */

	/*
	 * Information for management of parameter-change-driven rescanning
	 *
	 * extParam includes the paramIDs of all external PARAM_EXEC params
	 * affecting this plan node or its children.  setParam params from the
	 * node's initPlans are not included, but their extParams are.
	 *
	 * allParam includes all the extParam paramIDs, plus the IDs of local
	 * params that affect the node (i.e., the setParams of its initplans).
	 * These are _all_ the PARAM_EXEC params that affect this node.
	 */
	ExtParam []uint32 `json:"extParam"`
	AllParam []uint32 `json:"allParam"`
}

func (node Plan) MarshalJSON() ([]byte, error) {
	type PlanMarshalAlias Plan
	return json.Marshal(map[string]interface{}{
		"PLAN": (*PlanMarshalAlias)(&node),
	})
}

func (node *Plan) UnmarshalJSON(input []byte) (err error) {
	var fields map[string]json.RawMessage

	err = json.Unmarshal(input, &fields)
	if err != nil {
		return
	}

	if fields["startup_cost"] != nil {
		err = json.Unmarshal(fields["startup_cost"], &node.StartupCost)
		if err != nil {
			return
		}
	}

	if fields["total_cost"] != nil {
		err = json.Unmarshal(fields["total_cost"], &node.TotalCost)
		if err != nil {
			return
		}
	}

	if fields["plan_rows"] != nil {
		err = json.Unmarshal(fields["plan_rows"], &node.PlanRows)
		if err != nil {
			return
		}
	}

	if fields["plan_width"] != nil {
		err = json.Unmarshal(fields["plan_width"], &node.PlanWidth)
		if err != nil {
			return
		}
	}

	if fields["targetlist"] != nil {
		node.Targetlist, err = UnmarshalNodeArrayJSON(fields["targetlist"])
		if err != nil {
			return
		}
	}

	if fields["qual"] != nil {
		node.Qual, err = UnmarshalNodeArrayJSON(fields["qual"])
		if err != nil {
			return
		}
	}

	if fields["lefttree"] != nil {
		var nodePtr *Node
		nodePtr, err = UnmarshalNodePtrJSON(fields["lefttree"])
		if err != nil {
			return
		}
		if nodePtr != nil && *nodePtr != nil {
			val := (*nodePtr).(Plan)
			node.Lefttree = &val
		}
	}

	if fields["righttree"] != nil {
		var nodePtr *Node
		nodePtr, err = UnmarshalNodePtrJSON(fields["righttree"])
		if err != nil {
			return
		}
		if nodePtr != nil && *nodePtr != nil {
			val := (*nodePtr).(Plan)
			node.Righttree = &val
		}
	}

	if fields["initPlan"] != nil {
		node.InitPlan, err = UnmarshalNodeArrayJSON(fields["initPlan"])
		if err != nil {
			return
		}
	}

	if fields["extParam"] != nil {
		err = json.Unmarshal(fields["extParam"], &node.ExtParam)
		if err != nil {
			return
		}
	}

	if fields["allParam"] != nil {
		err = json.Unmarshal(fields["allParam"], &node.AllParam)
		if err != nil {
			return
		}
	}

	return
}
