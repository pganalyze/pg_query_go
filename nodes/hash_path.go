// Auto-generated from postgres/src/include/nodes/relation.h - DO NOT EDIT

package pg_query

import "encoding/json"

/*
 * A hashjoin path has these fields.
 *
 * The remarks above for mergeclauses apply for hashclauses as well.
 *
 * Hashjoin does not care what order its inputs appear in, so we have
 * no need for sortkeys.
 */
type HashPath struct {
	Jpath           JoinPath `json:"jpath"`
	PathHashclauses []Node   `json:"path_hashclauses"` /* join clauses used for hashing */
	NumBatches      int      `json:"num_batches"`      /* number of batches expected */
}

func (node HashPath) MarshalJSON() ([]byte, error) {
	type HashPathMarshalAlias HashPath
	return json.Marshal(map[string]interface{}{
		"HASHPATH": (*HashPathMarshalAlias)(&node),
	})
}

func (node *HashPath) UnmarshalJSON(input []byte) (err error) {
	var fields map[string]json.RawMessage

	err = json.Unmarshal(input, &fields)
	if err != nil {
		return
	}

	if fields["jpath"] != nil {
		err = json.Unmarshal(fields["jpath"], &node.Jpath)
		if err != nil {
			return
		}
	}

	if fields["path_hashclauses"] != nil {
		node.PathHashclauses, err = UnmarshalNodeArrayJSON(fields["path_hashclauses"])
		if err != nil {
			return
		}
	}

	if fields["num_batches"] != nil {
		err = json.Unmarshal(fields["num_batches"], &node.NumBatches)
		if err != nil {
			return
		}
	}

	return
}
