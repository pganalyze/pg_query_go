// Auto-generated from postgres/src/include/nodes/plannodes.h - DO NOT EDIT

package pg_query

import "encoding/json"

/* ----------------
 *	 Append node -
 *		Generate the concatenation of the results of sub-plans.
 * ----------------
 */
type Append struct {
	Plan        Plan   `json:"plan"`
	Appendplans []Node `json:"appendplans"`
}

func (node Append) MarshalJSON() ([]byte, error) {
	type AppendMarshalAlias Append
	return json.Marshal(map[string]interface{}{
		"APPEND": (*AppendMarshalAlias)(&node),
	})
}

func (node *Append) UnmarshalJSON(input []byte) (err error) {
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

	if fields["appendplans"] != nil {
		node.Appendplans, err = UnmarshalNodeArrayJSON(fields["appendplans"])
		if err != nil {
			return
		}
	}

	return
}
