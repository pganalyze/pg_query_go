// Auto-generated - DO NOT EDIT

package pg_query

type PlanRowMark struct {
	Rti       Index       `json:"rti"`       /* range table index of markable relation */
	Prti      Index       `json:"prti"`      /* range table index of parent relation */
	RowmarkId Index       `json:"rowmarkId"` /* unique identifier for resjunk columns */
	MarkType  RowMarkType `json:"markType"`  /* see enum above */
	NoWait    bool        `json:"noWait"`    /* NOWAIT option */
	IsParent  bool        `json:"isParent"`  /* true if this is a "dummy" parent entry */
}
