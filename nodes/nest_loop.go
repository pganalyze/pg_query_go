// Auto-generated from postgres/src/include/nodes/plannodes.h - DO NOT EDIT

package pg_query

import "encoding/json"

/* ----------------
 *		nest loop join node
 *
 * The nestParams list identifies any executor Params that must be passed
 * into execution of the inner subplan carrying values from the current row
 * of the outer subplan.  Currently we restrict these values to be simple
 * Vars, but perhaps someday that'd be worth relaxing.  (Note: during plan
 * creation, the paramval can actually be a PlaceHolderVar expression; but it
 * must be a Var with varno OUTER_VAR by the time it gets to the executor.)
 * ----------------
 */
type NestLoop struct {
	Join       Join   `json:"join"`
	NestParams []Node `json:"nestParams"` /* list of NestLoopParam nodes */
}

func (node NestLoop) MarshalJSON() ([]byte, error) {
	type NestLoopMarshalAlias NestLoop
	return json.Marshal(map[string]interface{}{
		"NESTLOOP": (*NestLoopMarshalAlias)(&node),
	})
}

func (node *NestLoop) UnmarshalJSON(input []byte) (err error) {
	var fields map[string]json.RawMessage

	err = json.Unmarshal(input, &fields)
	if err != nil {
		return
	}

	if fields["join"] != nil {
		err = json.Unmarshal(fields["join"], &node.Join)
		if err != nil {
			return
		}
	}

	if fields["nestParams"] != nil {
		node.NestParams, err = UnmarshalNodeArrayJSON(fields["nestParams"])
		if err != nil {
			return
		}
	}

	return
}
