// Auto-generated from postgres/src/include/nodes/relation.h - DO NOT EDIT

package pg_query

import "encoding/json"

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

func (node JoinCostWorkspace) MarshalJSON() ([]byte, error) {
	type JoinCostWorkspaceMarshalAlias JoinCostWorkspace
	return json.Marshal(map[string]interface{}{
		"JOINCOSTWORKSPACE": (*JoinCostWorkspaceMarshalAlias)(&node),
	})
}

func (node *JoinCostWorkspace) UnmarshalJSON(input []byte) (err error) {
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

	if fields["run_cost"] != nil {
		err = json.Unmarshal(fields["run_cost"], &node.RunCost)
		if err != nil {
			return
		}
	}

	if fields["inner_run_cost"] != nil {
		err = json.Unmarshal(fields["inner_run_cost"], &node.InnerRunCost)
		if err != nil {
			return
		}
	}

	if fields["inner_rescan_run_cost"] != nil {
		err = json.Unmarshal(fields["inner_rescan_run_cost"], &node.InnerRescanRunCost)
		if err != nil {
			return
		}
	}

	if fields["outer_rows"] != nil {
		err = json.Unmarshal(fields["outer_rows"], &node.OuterRows)
		if err != nil {
			return
		}
	}

	if fields["inner_rows"] != nil {
		err = json.Unmarshal(fields["inner_rows"], &node.InnerRows)
		if err != nil {
			return
		}
	}

	if fields["outer_skip_rows"] != nil {
		err = json.Unmarshal(fields["outer_skip_rows"], &node.OuterSkipRows)
		if err != nil {
			return
		}
	}

	if fields["inner_skip_rows"] != nil {
		err = json.Unmarshal(fields["inner_skip_rows"], &node.InnerSkipRows)
		if err != nil {
			return
		}
	}

	if fields["numbuckets"] != nil {
		err = json.Unmarshal(fields["numbuckets"], &node.Numbuckets)
		if err != nil {
			return
		}
	}

	if fields["numbatches"] != nil {
		err = json.Unmarshal(fields["numbatches"], &node.Numbatches)
		if err != nil {
			return
		}
	}

	return
}
