// Auto-generated from postgres/src/include/nodes/params.h - DO NOT EDIT

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
	var fields map[string]json.RawMessage

	err = json.Unmarshal(input, &fields)
	if err != nil {
		return
	}

	if fields["execPlan"] != nil {
		err = json.Unmarshal(fields["execPlan"], &node.ExecPlan)
		if err != nil {
			return
		}
	}

	if fields["value"] != nil {
		err = json.Unmarshal(fields["value"], &node.Value)
		if err != nil {
			return
		}
	}

	if fields["isnull"] != nil {
		err = json.Unmarshal(fields["isnull"], &node.Isnull)
		if err != nil {
			return
		}
	}

	return
}
