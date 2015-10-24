// Auto-generated - DO NOT EDIT

package pg_query

type MinMaxAggInfo struct {
	Aggfnoid  Oid          `json:"aggfnoid"`  /* pg_proc Oid of the aggregate */
	Aggsortop Oid          `json:"aggsortop"` /* Oid of its sort operator */
	Target    *Expr        `json:"target"`    /* expression we are aggregating on */
	Subroot   *PlannerInfo `json:"subroot"`   /* modified "root" for planning the subquery */
	Path      *Path        `json:"path"`      /* access path for subquery */
	Pathcost  Cost         `json:"pathcost"`  /* estimated cost to fetch first row */
	Param     *Param       `json:"param"`     /* param for subplan's output */
}
