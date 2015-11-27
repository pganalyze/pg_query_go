// Auto-generated from postgres/src/include/nodes/plannodes.h - DO NOT EDIT

package pg_query

import "encoding/json"

/* ----------------
 *		materialization node
 * ----------------
 */
type Material struct {
	Plan Plan `json:"plan"`
}

func (node Material) MarshalJSON() ([]byte, error) {
	type MaterialMarshalAlias Material
	return json.Marshal(map[string]interface{}{
		"MATERIAL": (*MaterialMarshalAlias)(&node),
	})
}

func (node *Material) UnmarshalJSON(input []byte) (err error) {
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

	return
}
