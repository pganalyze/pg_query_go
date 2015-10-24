// Auto-generated - DO NOT EDIT

package pg_query

type HashPath struct {
	Jpath           JoinPath `json:"jpath"`
	PathHashclauses []Node   `json:"path_hashclauses"` /* join clauses used for hashing */
	NumBatches      int      `json:"num_batches"`      /* number of batches expected */
}
