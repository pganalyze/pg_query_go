// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

type AlternativeSubPlan struct {
	Xpr      Expr   `json:"xpr"`
	Subplans []Node `json:"subplans"` /* SubPlan(s) with equivalent results */
}

func (node AlternativeSubPlan) MarshalJSON() ([]byte, error) {
	type AlternativeSubPlanMarshalAlias AlternativeSubPlan
	return json.Marshal(map[string]interface{}{
		"ALTERNATIVESUBPLAN": (*AlternativeSubPlanMarshalAlias)(&node),
	})
}

func (node *AlternativeSubPlan) UnmarshalJSON(input []byte) (err error) {
	var fields map[string]json.RawMessage

	err = json.Unmarshal(input, &fields)
	if err != nil {
		return
	}

	if fields["xpr"] != nil {
		err = json.Unmarshal(fields["xpr"], &node.Xpr)
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

	return
}
