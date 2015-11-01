// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

type LockingClause struct {
	LockedRels []Node             `json:"lockedRels"` /* FOR [KEY] UPDATE/SHARE relations */
	Strength   LockClauseStrength `json:"strength"`
	NoWait     bool               `json:"noWait"` /* NOWAIT option */
}

func (node LockingClause) MarshalJSON() ([]byte, error) {
	type LockingClauseMarshalAlias LockingClause
	return json.Marshal(map[string]interface{}{
		"LOCKINGCLAUSE": (*LockingClauseMarshalAlias)(&node),
	})
}

func (node *LockingClause) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}
