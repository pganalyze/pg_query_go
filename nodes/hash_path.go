// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

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
	err = UnmarshalNodeFieldJSON(input, node)
	return
}
