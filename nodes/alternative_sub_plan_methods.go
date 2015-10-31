// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

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

func (node AlternativeSubPlan) Deparse() string {
	panic("Not Implemented")
}
