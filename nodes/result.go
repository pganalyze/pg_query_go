// Auto-generated from postgres/src/include/nodes/plannodes.h - DO NOT EDIT

package pg_query

import "encoding/json"

/* ----------------
 *	 Result node -
 *		If no outer plan, evaluate a variable-free targetlist.
 *		If outer plan, return tuples from outer plan (after a level of
 *		projection as shown by targetlist).
 *
 * If resconstantqual isn't NULL, it represents a one-time qualification
 * test (i.e., one that doesn't depend on any variables from the outer plan,
 * so needs to be evaluated only once).
 * ----------------
 */
type Result struct {
	Plan            Plan `json:"plan"`
	Resconstantqual Node `json:"resconstantqual"`
}

func (node Result) MarshalJSON() ([]byte, error) {
	type ResultMarshalAlias Result
	return json.Marshal(map[string]interface{}{
		"RESULT": (*ResultMarshalAlias)(&node),
	})
}

func (node *Result) UnmarshalJSON(input []byte) (err error) {
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

	if fields["resconstantqual"] != nil {
		node.Resconstantqual, err = UnmarshalNodeJSON(fields["resconstantqual"])
		if err != nil {
			return
		}
	}

	return
}
