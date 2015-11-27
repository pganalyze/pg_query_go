// Auto-generated from postgres/src/include/nodes/plannodes.h - DO NOT EDIT

package pg_query

import "encoding/json"

type PlannedStmt struct {
	CommandType CmdType `json:"commandType"` /* select|insert|update|delete */

	QueryId uint32 `json:"queryId"` /* query identifier (copied from Query) */

	HasReturning bool `json:"hasReturning"` /* is it insert|update|delete RETURNING? */

	HasModifyingCte bool `json:"hasModifyingCTE"` /* has insert|update|delete in WITH? */

	CanSetTag bool `json:"canSetTag"` /* do I set the command result tag? */

	TransientPlan bool `json:"transientPlan"` /* redo plan when TransactionXmin changes? */

	PlanTree *Plan `json:"planTree"` /* tree of Plan nodes */

	Rtable []Node `json:"rtable"` /* list of RangeTblEntry nodes */

	/* rtable indexes of target relations for INSERT/UPDATE/DELETE */
	ResultRelations []Node `json:"resultRelations"` /* integer list of RT indexes, or NIL */

	UtilityStmt Node `json:"utilityStmt"` /* non-null if this is DECLARE CURSOR */

	Subplans []Node `json:"subplans"` /* Plan trees for SubPlan expressions */

	RewindPlanIds []uint32 `json:"rewindPlanIDs"` /* indices of subplans that require REWIND */

	RowMarks []Node `json:"rowMarks"` /* a list of PlanRowMark's */

	RelationOids []Node `json:"relationOids"` /* OIDs of relations the plan depends on */

	InvalItems []Node `json:"invalItems"` /* other dependencies, as PlanInvalItems */

	NParamExec int `json:"nParamExec"` /* number of PARAM_EXEC Params used */
}

func (node PlannedStmt) MarshalJSON() ([]byte, error) {
	type PlannedStmtMarshalAlias PlannedStmt
	return json.Marshal(map[string]interface{}{
		"PLANNEDSTMT": (*PlannedStmtMarshalAlias)(&node),
	})
}

func (node *PlannedStmt) UnmarshalJSON(input []byte) (err error) {
	var fields map[string]json.RawMessage

	err = json.Unmarshal(input, &fields)
	if err != nil {
		return
	}

	if fields["commandType"] != nil {
		err = json.Unmarshal(fields["commandType"], &node.CommandType)
		if err != nil {
			return
		}
	}

	if fields["queryId"] != nil {
		err = json.Unmarshal(fields["queryId"], &node.QueryId)
		if err != nil {
			return
		}
	}

	if fields["hasReturning"] != nil {
		err = json.Unmarshal(fields["hasReturning"], &node.HasReturning)
		if err != nil {
			return
		}
	}

	if fields["hasModifyingCTE"] != nil {
		err = json.Unmarshal(fields["hasModifyingCTE"], &node.HasModifyingCte)
		if err != nil {
			return
		}
	}

	if fields["canSetTag"] != nil {
		err = json.Unmarshal(fields["canSetTag"], &node.CanSetTag)
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

	if fields["planTree"] != nil {
		var nodePtr *Node
		nodePtr, err = UnmarshalNodePtrJSON(fields["planTree"])
		if err != nil {
			return
		}
		if nodePtr != nil && *nodePtr != nil {
			val := (*nodePtr).(Plan)
			node.PlanTree = &val
		}
	}

	if fields["rtable"] != nil {
		node.Rtable, err = UnmarshalNodeArrayJSON(fields["rtable"])
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

	if fields["utilityStmt"] != nil {
		node.UtilityStmt, err = UnmarshalNodeJSON(fields["utilityStmt"])
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

	if fields["rewindPlanIDs"] != nil {
		err = json.Unmarshal(fields["rewindPlanIDs"], &node.RewindPlanIds)
		if err != nil {
			return
		}
	}

	if fields["rowMarks"] != nil {
		node.RowMarks, err = UnmarshalNodeArrayJSON(fields["rowMarks"])
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

	return
}
