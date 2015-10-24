// Auto-generated - DO NOT EDIT

package pg_query

type JoinCostWorkspace struct {
	/* Preliminary cost estimates --- must not be larger than final ones! */
	StartupCost Cost `json:"startup_cost"` /* cost expended before fetching any tuples */
	TotalCost   Cost `json:"total_cost"`   /* total cost (assuming all tuples fetched) */

	/* Fields below here should be treated as private to costsize.c */
	RunCost Cost `json:"run_cost"` /* non-startup cost components */

	/* private for cost_nestloop code */
	InnerRunCost       Cost `json:"inner_run_cost"` /* also used by cost_mergejoin code */
	InnerRescanRunCost Cost `json:"inner_rescan_run_cost"`

	/* private for cost_mergejoin code */
	OuterRows     float64 `json:"outer_rows"`
	InnerRows     float64 `json:"inner_rows"`
	OuterSkipRows float64 `json:"outer_skip_rows"`
	InnerSkipRows float64 `json:"inner_skip_rows"`

	/* private for cost_hashjoin code */
	Numbuckets int `json:"numbuckets"`
	Numbatches int `json:"numbatches"`
}
