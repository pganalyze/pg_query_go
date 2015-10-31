// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

func (node AggClauseCosts) MarshalJSON() ([]byte, error) {
	type AggClauseCostsMarshalAlias AggClauseCosts
	return json.Marshal(map[string]interface{}{
		"AGGCLAUSECOSTS": (*AggClauseCostsMarshalAlias)(&node),
	})
}

func (node *AggClauseCosts) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}

func (node AggClauseCosts) Deparse() string {
	panic("Not Implemented")
}
