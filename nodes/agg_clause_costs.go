// Auto-generated from postgres/src/include/nodes/relation.h - DO NOT EDIT

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
	var fields map[string]json.RawMessage

	err = json.Unmarshal(input, &fields)
	if err != nil {
		return
	}

	if fields["numAggs"] != nil {
		err = json.Unmarshal(fields["numAggs"], &node.NumAggs)
		if err != nil {
			return
		}
	}

	if fields["numOrderedAggs"] != nil {
		err = json.Unmarshal(fields["numOrderedAggs"], &node.NumOrderedAggs)
		if err != nil {
			return
		}
	}

	if fields["transCost"] != nil {
		err = json.Unmarshal(fields["transCost"], &node.TransCost)
		if err != nil {
			return
		}
	}

	if fields["finalCost"] != nil {
		err = json.Unmarshal(fields["finalCost"], &node.FinalCost)
		if err != nil {
			return
		}
	}

	if fields["transitionSpace"] != nil {
		err = json.Unmarshal(fields["transitionSpace"], &node.TransitionSpace)
		if err != nil {
			return
		}
	}

	return
}
