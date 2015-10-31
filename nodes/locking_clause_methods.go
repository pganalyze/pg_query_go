// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

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

func (node LockingClause) Deparse() string {
	panic("Not Implemented")
}
