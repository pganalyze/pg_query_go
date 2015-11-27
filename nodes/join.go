// Auto-generated from postgres/src/include/nodes/plannodes.h - DO NOT EDIT

package pg_query

import "encoding/json"

/* ----------------
 *		Join node
 *
 * jointype:	rule for joining tuples from left and right subtrees
 * joinqual:	qual conditions that came from JOIN/ON or JOIN/USING
 *				(plan.qual contains conditions that came from WHERE)
 *
 * When jointype is INNER, joinqual and plan.qual are semantically
 * interchangeable.  For OUTER jointypes, the two are *not* interchangeable;
 * only joinqual is used to determine whether a match has been found for
 * the purpose of deciding whether to generate null-extended tuples.
 * (But plan.qual is still applied before actually returning a tuple.)
 * For an outer join, only joinquals are allowed to be used as the merge
 * or hash condition of a merge or hash join.
 * ----------------
 */
type Join struct {
	Plan     Plan     `json:"plan"`
	Jointype JoinType `json:"jointype"`
	Joinqual []Node   `json:"joinqual"` /* JOIN quals (in addition to plan.qual) */
}

func (node Join) MarshalJSON() ([]byte, error) {
	type JoinMarshalAlias Join
	return json.Marshal(map[string]interface{}{
		"JOIN": (*JoinMarshalAlias)(&node),
	})
}

func (node *Join) UnmarshalJSON(input []byte) (err error) {
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

	if fields["jointype"] != nil {
		err = json.Unmarshal(fields["jointype"], &node.Jointype)
		if err != nil {
			return
		}
	}

	if fields["joinqual"] != nil {
		node.Joinqual, err = UnmarshalNodeArrayJSON(fields["joinqual"])
		if err != nil {
			return
		}
	}

	return
}
