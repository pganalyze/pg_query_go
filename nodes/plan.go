// Auto-generated - DO NOT EDIT

package pg_query

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
