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
	err = UnmarshalNodeFieldJSON(input, node)
	return
}
