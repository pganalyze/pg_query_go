// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

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

func (node QualCost) Deparse() string {
	panic("Not Implemented")
}
