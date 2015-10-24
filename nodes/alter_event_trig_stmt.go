// Auto-generated - DO NOT EDIT

package pg_query

type AlterEventTrigStmt struct {
	Trigname  *string `json:"trigname"`  /* TRIGGER's name */
	Tgenabled byte    `json:"tgenabled"` /* trigger's firing configuration WRT
	 * session_replication_role */
}
