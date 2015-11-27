// Auto-generated from postgres/src/include/nodes/plannodes.h - DO NOT EDIT

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
	var fields map[string]json.RawMessage

	err = json.Unmarshal(input, &fields)
	if err != nil {
		return
	}

	if fields["plan"] != nil {
		err = json.Unmarshal(fields["plan"], &node.Plan)
		if err != nil {
			return
		}
	}

	if fields["rowMarks"] != nil {
		node.RowMarks, err = UnmarshalNodeArrayJSON(fields["rowMarks"])
		if err != nil {
			return
		}
	}

	if fields["epqParam"] != nil {
		err = json.Unmarshal(fields["epqParam"], &node.EpqParam)
		if err != nil {
			return
		}
	}

	return
}
