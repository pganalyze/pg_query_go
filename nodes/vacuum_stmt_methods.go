// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

func (node VacuumStmt) MarshalJSON() ([]byte, error) {
	type VacuumStmtMarshalAlias VacuumStmt
	return json.Marshal(map[string]interface{}{
		"VACUUM": (*VacuumStmtMarshalAlias)(&node),
	})
}

func (node *VacuumStmt) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}

func (node VacuumStmt) Deparse() string {
	panic("Not Implemented")
}
