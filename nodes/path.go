// Auto-generated - DO NOT EDIT

package pg_query

type Path struct {
	Parent    *RelOptInfo    `json:"parent"`     /* the relation this path can build */
	ParamInfo *ParamPathInfo `json:"param_info"` /* parameterization info, or NULL if none */

	/* estimated size/costs for path (see costsize.c for more info) */
	Rows        float64 `json:"rows"`         /* estimated number of result tuples */
	StartupCost Cost    `json:"startup_cost"` /* cost expended before fetching any tuples */
	TotalCost   Cost    `json:"total_cost"`   /* total cost (assuming all tuples fetched) */

	Pathkeys []Node `json:"pathkeys"` /* sort ordering of path's output */
	/* pathkeys is a List of PathKey nodes; see above */
}
