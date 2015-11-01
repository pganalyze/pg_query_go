// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

type MinMaxAggInfo struct {
	Aggfnoid  Oid          `json:"aggfnoid"`  /* pg_proc Oid of the aggregate */
	Aggsortop Oid          `json:"aggsortop"` /* Oid of its sort operator */
	Target    *Expr        `json:"target"`    /* expression we are aggregating on */
	Subroot   *PlannerInfo `json:"subroot"`   /* modified "root" for planning the subquery */
	Path      *Path        `json:"path"`      /* access path for subquery */
	Pathcost  Cost         `json:"pathcost"`  /* estimated cost to fetch first row */
	Param     *Param       `json:"param"`     /* param for subplan's output */
}

func (node MinMaxAggInfo) MarshalJSON() ([]byte, error) {
	type MinMaxAggInfoMarshalAlias MinMaxAggInfo
	return json.Marshal(map[string]interface{}{
		"MINMAXAGGINFO": (*MinMaxAggInfoMarshalAlias)(&node),
	})
}

func (node *MinMaxAggInfo) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}
