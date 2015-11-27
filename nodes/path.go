// Auto-generated from postgres/src/include/nodes/relation.h - DO NOT EDIT

package pg_query

import "encoding/json"

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

func (node Path) MarshalJSON() ([]byte, error) {
	type PathMarshalAlias Path
	return json.Marshal(map[string]interface{}{
		"PATH": (*PathMarshalAlias)(&node),
	})
}

func (node *Path) UnmarshalJSON(input []byte) (err error) {
	var fields map[string]json.RawMessage

	err = json.Unmarshal(input, &fields)
	if err != nil {
		return
	}

	if fields["parent"] != nil {
		var nodePtr *Node
		nodePtr, err = UnmarshalNodePtrJSON(fields["parent"])
		if err != nil {
			return
		}
		if nodePtr != nil && *nodePtr != nil {
			val := (*nodePtr).(RelOptInfo)
			node.Parent = &val
		}
	}

	if fields["param_info"] != nil {
		var nodePtr *Node
		nodePtr, err = UnmarshalNodePtrJSON(fields["param_info"])
		if err != nil {
			return
		}
		if nodePtr != nil && *nodePtr != nil {
			val := (*nodePtr).(ParamPathInfo)
			node.ParamInfo = &val
		}
	}

	if fields["rows"] != nil {
		err = json.Unmarshal(fields["rows"], &node.Rows)
		if err != nil {
			return
		}
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

	if fields["pathkeys"] != nil {
		node.Pathkeys, err = UnmarshalNodeArrayJSON(fields["pathkeys"])
		if err != nil {
			return
		}
	}

	return
}
