// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

type QualCost struct {
	Startup  Cost `json:"startup"`   /* one-time cost */
	PerTuple Cost `json:"per_tuple"` /* per-evaluation cost */
}

func (node QualCost) MarshalJSON() ([]byte, error) {
	type QualCostMarshalAlias QualCost
	return json.Marshal(map[string]interface{}{
		"QUALCOST": (*QualCostMarshalAlias)(&node),
	})
}

func (node *QualCost) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}
