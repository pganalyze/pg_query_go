// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

type VacuumStmt struct {
	Options               int `json:"options"`                  /* OR of VacuumOption flags */
	FreezeMinAge          int `json:"freeze_min_age"`           /* min freeze age, or -1 to use default */
	FreezeTableAge        int `json:"freeze_table_age"`         /* age at which to scan whole table */
	MultixactFreezeMinAge int `json:"multixact_freeze_min_age"` /* min multixact freeze age,
	 * or -1 to use default */
	MultixactFreezeTableAge int `json:"multixact_freeze_table_age"` /* multixact age at which to
	 * scan whole table */
	Relation *RangeVar `json:"relation"` /* single table to process, or NULL */
	VaCols   []Node    `json:"va_cols"`  /* list of column names, or NIL for all */
}

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
