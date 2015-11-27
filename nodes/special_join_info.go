// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

type SpecialJoinInfo struct {
	MinLefthand     []uint32 `json:"min_lefthand"`      /* base relids in minimum LHS for join */
	MinRighthand    []uint32 `json:"min_righthand"`     /* base relids in minimum RHS for join */
	SynLefthand     []uint32 `json:"syn_lefthand"`      /* base relids syntactically within LHS */
	SynRighthand    []uint32 `json:"syn_righthand"`     /* base relids syntactically within RHS */
	Jointype        JoinType `json:"jointype"`          /* always INNER, LEFT, FULL, SEMI, or ANTI */
	LhsStrict       bool     `json:"lhs_strict"`        /* joinclause is strict for some LHS rel */
	DelayUpperJoins bool     `json:"delay_upper_joins"` /* can't commute with upper RHS */
	JoinQuals       []Node   `json:"join_quals"`        /* join quals, in implicit-AND list format */
}

func (node SpecialJoinInfo) MarshalJSON() ([]byte, error) {
	type SpecialJoinInfoMarshalAlias SpecialJoinInfo
	return json.Marshal(map[string]interface{}{
		"SPECIALJOININFO": (*SpecialJoinInfoMarshalAlias)(&node),
	})
}

func (node *SpecialJoinInfo) UnmarshalJSON(input []byte) (err error) {
	var fields map[string]json.RawMessage

	err = json.Unmarshal(input, &fields)
	if err != nil {
		return
	}

	if fields["min_lefthand"] != nil {
		err = json.Unmarshal(fields["min_lefthand"], &node.MinLefthand)
		if err != nil {
			return
		}
	}

	if fields["min_righthand"] != nil {
		err = json.Unmarshal(fields["min_righthand"], &node.MinRighthand)
		if err != nil {
			return
		}
	}

	if fields["syn_lefthand"] != nil {
		err = json.Unmarshal(fields["syn_lefthand"], &node.SynLefthand)
		if err != nil {
			return
		}
	}

	if fields["syn_righthand"] != nil {
		err = json.Unmarshal(fields["syn_righthand"], &node.SynRighthand)
		if err != nil {
			return
		}
	}

	if fields["jointype"] != nil {
		err = json.Unmarshal(fields["jointype"], &node.Jointype)
		if err != nil {
			return
		}
	}

	if fields["lhs_strict"] != nil {
		err = json.Unmarshal(fields["lhs_strict"], &node.LhsStrict)
		if err != nil {
			return
		}
	}

	if fields["delay_upper_joins"] != nil {
		err = json.Unmarshal(fields["delay_upper_joins"], &node.DelayUpperJoins)
		if err != nil {
			return
		}
	}

	if fields["join_quals"] != nil {
		node.JoinQuals, err = UnmarshalNodeArrayJSON(fields["join_quals"])
		if err != nil {
			return
		}
	}

	return
}
