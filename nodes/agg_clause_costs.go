// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

type AggClauseCosts struct {
	NumAggs         int      `json:"numAggs"`         /* total number of aggregate functions */
	NumOrderedAggs  int      `json:"numOrderedAggs"`  /* number w/ DISTINCT/ORDER BY/WITHIN GROUP */
	TransCost       QualCost `json:"transCost"`       /* total per-input-row execution costs */
	FinalCost       Cost     `json:"finalCost"`       /* total per-aggregated-row costs */
	TransitionSpace uintptr  `json:"transitionSpace"` /* space for pass-by-ref transition data */
}

func (node AggClauseCosts) MarshalJSON() ([]byte, error) {
	type AggClauseCostsMarshalAlias AggClauseCosts
	return json.Marshal(map[string]interface{}{
		"AGGCLAUSECOSTS": (*AggClauseCostsMarshalAlias)(&node),
	})
}

func (node *AggClauseCosts) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}
