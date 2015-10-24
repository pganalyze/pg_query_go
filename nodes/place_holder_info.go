// Auto-generated - DO NOT EDIT

package pg_query

type PlaceHolderInfo struct {
	Phid      Index           `json:"phid"`       /* ID for PH (unique within planner run) */
	PhVar     *PlaceHolderVar `json:"ph_var"`     /* copy of PlaceHolderVar tree */
	PhEvalAt  []uint32        `json:"ph_eval_at"` /* lowest level we can evaluate value at */
	PhLateral []uint32        `json:"ph_lateral"` /* relids of contained lateral refs, if any */
	PhNeeded  []uint32        `json:"ph_needed"`  /* highest level the value is needed at */
	PhWidth   int32           `json:"ph_width"`   /* estimated attribute width */
}
