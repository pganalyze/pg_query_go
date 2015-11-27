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
	var fields map[string]json.RawMessage

	err = json.Unmarshal(input, &fields)
	if err != nil {
		return
	}

	if fields["scan"] != nil {
		err = json.Unmarshal(fields["scan"], &node.Scan)
		if err != nil {
			return
		}
	}

	if fields["ctePlanId"] != nil {
		err = json.Unmarshal(fields["ctePlanId"], &node.CtePlanId)
		if err != nil {
			return
		}
	}

	if fields["cteParam"] != nil {
		err = json.Unmarshal(fields["cteParam"], &node.CteParam)
		if err != nil {
			return
		}
	}

	return
}
