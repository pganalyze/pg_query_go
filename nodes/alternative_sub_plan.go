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
	err = UnmarshalNodeFieldJSON(input, node)
	return
}
