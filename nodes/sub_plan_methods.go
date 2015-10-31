// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

func (node SubPlan) MarshalJSON() ([]byte, error) {
	type SubPlanMarshalAlias SubPlan
	return json.Marshal(map[string]interface{}{
		"SUBPLAN": (*SubPlanMarshalAlias)(&node),
	})
}

func (node *SubPlan) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}

func (node SubPlan) Deparse() string {
	panic("Not Implemented")
}
