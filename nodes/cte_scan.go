// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

type CteScan struct {
	Scan      Scan `json:"scan"`
	CtePlanId int  `json:"ctePlanId"` /* ID of init SubPlan for CTE */
	CteParam  int  `json:"cteParam"`  /* ID of Param representing CTE output */
}

func (node CteScan) MarshalJSON() ([]byte, error) {
	type CteScanMarshalAlias CteScan
	return json.Marshal(map[string]interface{}{
		"CTESCAN": (*CteScanMarshalAlias)(&node),
	})
}

func (node *CteScan) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}
