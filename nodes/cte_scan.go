// Auto-generated - DO NOT EDIT

package pg_query

type CteScan struct {
	Scan      Scan `json:"scan"`
	CtePlanId int  `json:"ctePlanId"` /* ID of init SubPlan for CTE */
	CteParam  int  `json:"cteParam"`  /* ID of Param representing CTE output */
}
