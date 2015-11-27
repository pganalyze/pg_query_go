// Auto-generated from postgres/src/include/nodes/relation.h - DO NOT EDIT

package pg_query

import "encoding/json"

/*
 * For each distinct placeholder expression generated during planning, we
 * store a PlaceHolderInfo node in the PlannerInfo node's placeholder_list.
 * This stores info that is needed centrally rather than in each copy of the
 * PlaceHolderVar.  The phid fields identify which PlaceHolderInfo goes with
 * each PlaceHolderVar.  Note that phid is unique throughout a planner run,
 * not just within a query level --- this is so that we need not reassign ID's
 * when pulling a subquery into its parent.
 *
 * The idea is to evaluate the expression at (only) the ph_eval_at join level,
 * then allow it to bubble up like a Var until the ph_needed join level.
 * ph_needed has the same definition as attr_needed for a regular Var.
 *
 * The PlaceHolderVar's expression might contain LATERAL references to vars
 * coming from outside its syntactic scope.  If so, those rels are *not*
 * included in ph_eval_at, but they are recorded in ph_lateral.
 *
 * Notice that when ph_eval_at is a join rather than a single baserel, the
 * PlaceHolderInfo may create constraints on join order: the ph_eval_at join
 * has to be formed below any outer joins that should null the PlaceHolderVar.
 *
 * We create a PlaceHolderInfo only after determining that the PlaceHolderVar
 * is actually referenced in the plan tree, so that unreferenced placeholders
 * don't result in unnecessary constraints on join order.
 */
type PlaceHolderInfo struct {
	Phid      Index           `json:"phid"`       /* ID for PH (unique within planner run) */
	PhVar     *PlaceHolderVar `json:"ph_var"`     /* copy of PlaceHolderVar tree */
	PhEvalAt  []uint32        `json:"ph_eval_at"` /* lowest level we can evaluate value at */
	PhLateral []uint32        `json:"ph_lateral"` /* relids of contained lateral refs, if any */
	PhNeeded  []uint32        `json:"ph_needed"`  /* highest level the value is needed at */
	PhWidth   int32           `json:"ph_width"`   /* estimated attribute width */
}

func (node PlaceHolderInfo) MarshalJSON() ([]byte, error) {
	type PlaceHolderInfoMarshalAlias PlaceHolderInfo
	return json.Marshal(map[string]interface{}{
		"PLACEHOLDERINFO": (*PlaceHolderInfoMarshalAlias)(&node),
	})
}

func (node *PlaceHolderInfo) UnmarshalJSON(input []byte) (err error) {
	var fields map[string]json.RawMessage

	err = json.Unmarshal(input, &fields)
	if err != nil {
		return
	}

	if fields["phid"] != nil {
		err = json.Unmarshal(fields["phid"], &node.Phid)
		if err != nil {
			return
		}
	}

	if fields["ph_var"] != nil {
		var nodePtr *Node
		nodePtr, err = UnmarshalNodePtrJSON(fields["ph_var"])
		if err != nil {
			return
		}
		if nodePtr != nil && *nodePtr != nil {
			val := (*nodePtr).(PlaceHolderVar)
			node.PhVar = &val
		}
	}

	if fields["ph_eval_at"] != nil {
		err = json.Unmarshal(fields["ph_eval_at"], &node.PhEvalAt)
		if err != nil {
			return
		}
	}

	if fields["ph_lateral"] != nil {
		err = json.Unmarshal(fields["ph_lateral"], &node.PhLateral)
		if err != nil {
			return
		}
	}

	if fields["ph_needed"] != nil {
		err = json.Unmarshal(fields["ph_needed"], &node.PhNeeded)
		if err != nil {
			return
		}
	}

	if fields["ph_width"] != nil {
		err = json.Unmarshal(fields["ph_width"], &node.PhWidth)
		if err != nil {
			return
		}
	}

	return
}
