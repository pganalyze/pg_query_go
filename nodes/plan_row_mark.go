// Auto-generated from postgres/src/include/nodes/plannodes.h - DO NOT EDIT

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
	var fields map[string]json.RawMessage

	err = json.Unmarshal(input, &fields)
	if err != nil {
		return
	}

	if fields["rti"] != nil {
		err = json.Unmarshal(fields["rti"], &node.Rti)
		if err != nil {
			return
		}
	}

	if fields["prti"] != nil {
		err = json.Unmarshal(fields["prti"], &node.Prti)
		if err != nil {
			return
		}
	}

	if fields["rowmarkId"] != nil {
		err = json.Unmarshal(fields["rowmarkId"], &node.RowmarkId)
		if err != nil {
			return
		}
	}

	if fields["markType"] != nil {
		err = json.Unmarshal(fields["markType"], &node.MarkType)
		if err != nil {
			return
		}
	}

	if fields["noWait"] != nil {
		err = json.Unmarshal(fields["noWait"], &node.NoWait)
		if err != nil {
			return
		}
	}

	if fields["isParent"] != nil {
		err = json.Unmarshal(fields["isParent"], &node.IsParent)
		if err != nil {
			return
		}
	}

	return
}
