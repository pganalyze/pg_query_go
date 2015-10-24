// Auto-generated - DO NOT EDIT

package pg_query

type LockRows struct {
	Plan     Plan   `json:"plan"`
	RowMarks []Node `json:"rowMarks"` /* a list of PlanRowMark's */
	EpqParam int    `json:"epqParam"` /* ID of Param for EvalPlanQual re-eval */
}
