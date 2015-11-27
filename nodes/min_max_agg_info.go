// Auto-generated from postgres/src/include/nodes/relation.h - DO NOT EDIT

package pg_query

import "encoding/json"

/*
 * For each potentially index-optimizable MIN/MAX aggregate function,
 * root->minmax_aggs stores a MinMaxAggInfo describing it.
 */
type MinMaxAggInfo struct {
	Aggfnoid  Oid          `json:"aggfnoid"`  /* pg_proc Oid of the aggregate */
	Aggsortop Oid          `json:"aggsortop"` /* Oid of its sort operator */
	Target    *Expr        `json:"target"`    /* expression we are aggregating on */
	Subroot   *PlannerInfo `json:"subroot"`   /* modified "root" for planning the subquery */
	Path      *Path        `json:"path"`      /* access path for subquery */
	Pathcost  Cost         `json:"pathcost"`  /* estimated cost to fetch first row */
	Param     *Param       `json:"param"`     /* param for subplan's output */
}

func (node MinMaxAggInfo) MarshalJSON() ([]byte, error) {
	type MinMaxAggInfoMarshalAlias MinMaxAggInfo
	return json.Marshal(map[string]interface{}{
		"MINMAXAGGINFO": (*MinMaxAggInfoMarshalAlias)(&node),
	})
}

func (node *MinMaxAggInfo) UnmarshalJSON(input []byte) (err error) {
	var fields map[string]json.RawMessage

	err = json.Unmarshal(input, &fields)
	if err != nil {
		return
	}

	if fields["aggfnoid"] != nil {
		err = json.Unmarshal(fields["aggfnoid"], &node.Aggfnoid)
		if err != nil {
			return
		}
	}

	if fields["aggsortop"] != nil {
		err = json.Unmarshal(fields["aggsortop"], &node.Aggsortop)
		if err != nil {
			return
		}
	}

	if fields["target"] != nil {
		var nodePtr *Node
		nodePtr, err = UnmarshalNodePtrJSON(fields["target"])
		if err != nil {
			return
		}
		if nodePtr != nil && *nodePtr != nil {
			val := (*nodePtr).(Expr)
			node.Target = &val
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

	if fields["path"] != nil {
		var nodePtr *Node
		nodePtr, err = UnmarshalNodePtrJSON(fields["path"])
		if err != nil {
			return
		}
		if nodePtr != nil && *nodePtr != nil {
			val := (*nodePtr).(Path)
			node.Path = &val
		}
	}

	if fields["pathcost"] != nil {
		err = json.Unmarshal(fields["pathcost"], &node.Pathcost)
		if err != nil {
			return
		}
	}

	if fields["param"] != nil {
		var nodePtr *Node
		nodePtr, err = UnmarshalNodePtrJSON(fields["param"])
		if err != nil {
			return
		}
		if nodePtr != nil && *nodePtr != nil {
			val := (*nodePtr).(Param)
			node.Param = &val
		}
	}

	return
}
