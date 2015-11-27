// Auto-generated from postgres/src/include/nodes/relation.h - DO NOT EDIT

package pg_query

import "encoding/json"

/*----------
 * PlannerGlobal
 *		Global information for planning/optimization
 *
 * PlannerGlobal holds state for an entire planner invocation; this state
 * is shared across all levels of sub-Queries that exist in the command being
 * planned.
 *----------
 */
type PlannerGlobal struct {
	BoundParams ParamListInfo `json:"boundParams"` /* Param values provided to planner() */

	Subplans []Node `json:"subplans"` /* Plans for SubPlan nodes */

	Subroots []Node `json:"subroots"` /* PlannerInfos for SubPlan nodes */

	RewindPlanIds []uint32 `json:"rewindPlanIDs"` /* indices of subplans that require REWIND */

	Finalrtable []Node `json:"finalrtable"` /* "flat" rangetable for executor */

	Finalrowmarks []Node `json:"finalrowmarks"` /* "flat" list of PlanRowMarks */

	ResultRelations []Node `json:"resultRelations"` /* "flat" list of integer RT indexes */

	RelationOids []Node `json:"relationOids"` /* OIDs of relations the plan depends on */

	InvalItems []Node `json:"invalItems"` /* other dependencies, as PlanInvalItems */

	NParamExec int `json:"nParamExec"` /* number of PARAM_EXEC Params used */

	LastPhid Index `json:"lastPHId"` /* highest PlaceHolderVar ID assigned */

	LastRowMarkId Index `json:"lastRowMarkId"` /* highest PlanRowMark ID assigned */

	TransientPlan bool `json:"transientPlan"` /* redo plan when TransactionXmin changes? */
}

func (node PlannerGlobal) MarshalJSON() ([]byte, error) {
	type PlannerGlobalMarshalAlias PlannerGlobal
	return json.Marshal(map[string]interface{}{
		"PLANNERGLOBAL": (*PlannerGlobalMarshalAlias)(&node),
	})
}

func (node *PlannerGlobal) UnmarshalJSON(input []byte) (err error) {
	var fields map[string]json.RawMessage

	err = json.Unmarshal(input, &fields)
	if err != nil {
		return
	}

	if fields["boundParams"] != nil {
		err = json.Unmarshal(fields["boundParams"], &node.BoundParams)
		if err != nil {
			return
		}
	}

	if fields["subplans"] != nil {
		node.Subplans, err = UnmarshalNodeArrayJSON(fields["subplans"])
		if err != nil {
			return
		}
	}

	if fields["subroots"] != nil {
		node.Subroots, err = UnmarshalNodeArrayJSON(fields["subroots"])
		if err != nil {
			return
		}
	}

	if fields["rewindPlanIDs"] != nil {
		err = json.Unmarshal(fields["rewindPlanIDs"], &node.RewindPlanIds)
		if err != nil {
			return
		}
	}

	if fields["finalrtable"] != nil {
		node.Finalrtable, err = UnmarshalNodeArrayJSON(fields["finalrtable"])
		if err != nil {
			return
		}
	}

	if fields["finalrowmarks"] != nil {
		node.Finalrowmarks, err = UnmarshalNodeArrayJSON(fields["finalrowmarks"])
		if err != nil {
			return
		}
	}

	if fields["resultRelations"] != nil {
		node.ResultRelations, err = UnmarshalNodeArrayJSON(fields["resultRelations"])
		if err != nil {
			return
		}
	}

	if fields["relationOids"] != nil {
		node.RelationOids, err = UnmarshalNodeArrayJSON(fields["relationOids"])
		if err != nil {
			return
		}
	}

	if fields["invalItems"] != nil {
		node.InvalItems, err = UnmarshalNodeArrayJSON(fields["invalItems"])
		if err != nil {
			return
		}
	}

	if fields["nParamExec"] != nil {
		err = json.Unmarshal(fields["nParamExec"], &node.NParamExec)
		if err != nil {
			return
		}
	}

	if fields["lastPHId"] != nil {
		err = json.Unmarshal(fields["lastPHId"], &node.LastPhid)
		if err != nil {
			return
		}
	}

	if fields["lastRowMarkId"] != nil {
		err = json.Unmarshal(fields["lastRowMarkId"], &node.LastRowMarkId)
		if err != nil {
			return
		}
	}

	if fields["transientPlan"] != nil {
		err = json.Unmarshal(fields["transientPlan"], &node.TransientPlan)
		if err != nil {
			return
		}
	}

	return
}
