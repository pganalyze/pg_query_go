// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

type LockRows struct {
	Plan     Plan   `json:"plan"`
	RowMarks []Node `json:"rowMarks"` /* a list of PlanRowMark's */
	EpqParam int    `json:"epqParam"` /* ID of Param for EvalPlanQual re-eval */
}

func (node LockRows) MarshalJSON() ([]byte, error) {
	type LockRowsMarshalAlias LockRows
	return json.Marshal(map[string]interface{}{
		"LOCKROWS": (*LockRowsMarshalAlias)(&node),
	})
}

func (node *LockRows) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}
