// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

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
