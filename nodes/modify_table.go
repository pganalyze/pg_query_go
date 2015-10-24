// Auto-generated - DO NOT EDIT

package pg_query

type ModifyTable struct {
	Plan                 Plan    `json:"plan"`
	Operation            CmdType `json:"operation"`            /* INSERT, UPDATE, or DELETE */
	CanSetTag            bool    `json:"canSetTag"`            /* do we set the command tag/es_processed? */
	ResultRelations      []Node  `json:"resultRelations"`      /* integer list of RT indexes */
	ResultRelIndex       int     `json:"resultRelIndex"`       /* index of first resultRel in plan's list */
	Plans                []Node  `json:"plans"`                /* plan(s) producing source data */
	WithCheckOptionLists []Node  `json:"withCheckOptionLists"` /* per-target-table WCO lists */
	ReturningLists       []Node  `json:"returningLists"`       /* per-target-table RETURNING tlists */
	FdwPrivLists         []Node  `json:"fdwPrivLists"`         /* per-target-table FDW private data lists */
	RowMarks             []Node  `json:"rowMarks"`             /* PlanRowMarks (non-locking only) */
	EpqParam             int     `json:"epqParam"`             /* ID of Param for EvalPlanQual re-eval */
}
