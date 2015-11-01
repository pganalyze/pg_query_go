// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

type RowMarkClause struct {
	Rti        Index              `json:"rti"` /* range table index of target relation */
	Strength   LockClauseStrength `json:"strength"`
	NoWait     bool               `json:"noWait"`     /* NOWAIT option */
	PushedDown bool               `json:"pushedDown"` /* pushed down from higher query level? */
}

func (node RowMarkClause) MarshalJSON() ([]byte, error) {
	type RowMarkClauseMarshalAlias RowMarkClause
	return json.Marshal(map[string]interface{}{
		"ROWMARKCLAUSE": (*RowMarkClauseMarshalAlias)(&node),
	})
}

func (node *RowMarkClause) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}
