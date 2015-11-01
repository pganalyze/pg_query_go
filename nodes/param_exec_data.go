// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

type ParamExecData struct {
	ExecPlan interface{} `json:"execPlan"` /* should be "SubPlanState *" */
	Value    Datum       `json:"value"`
	Isnull   bool        `json:"isnull"`
}

func (node ParamExecData) MarshalJSON() ([]byte, error) {
	type ParamExecDataMarshalAlias ParamExecData
	return json.Marshal(map[string]interface{}{
		"PARAMEXECDATA": (*ParamExecDataMarshalAlias)(&node),
	})
}

func (node *ParamExecData) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}
