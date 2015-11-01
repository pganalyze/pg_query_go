// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

type PlanRowMark struct {
	Rti       Index       `json:"rti"`       /* range table index of markable relation */
	Prti      Index       `json:"prti"`      /* range table index of parent relation */
	RowmarkId Index       `json:"rowmarkId"` /* unique identifier for resjunk columns */
	MarkType  RowMarkType `json:"markType"`  /* see enum above */
	NoWait    bool        `json:"noWait"`    /* NOWAIT option */
	IsParent  bool        `json:"isParent"`  /* true if this is a "dummy" parent entry */
}

func (node PlanRowMark) MarshalJSON() ([]byte, error) {
	type PlanRowMarkMarshalAlias PlanRowMark
	return json.Marshal(map[string]interface{}{
		"PLANROWMARK": (*PlanRowMarkMarshalAlias)(&node),
	})
}

func (node *PlanRowMark) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}
