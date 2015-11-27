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
	var fields map[string]json.RawMessage

	err = json.Unmarshal(input, &fields)
	if err != nil {
		return
	}

	if fields["startup"] != nil {
		err = json.Unmarshal(fields["startup"], &node.Startup)
		if err != nil {
			return
		}
	}

	if fields["per_tuple"] != nil {
		err = json.Unmarshal(fields["per_tuple"], &node.PerTuple)
		if err != nil {
			return
		}
	}

	return
}
