// Auto-generated - DO NOT EDIT

package pg_query

type RowMarkClause struct {
	Rti        Index              `json:"rti"` /* range table index of target relation */
	Strength   LockClauseStrength `json:"strength"`
	NoWait     bool               `json:"noWait"`     /* NOWAIT option */
	PushedDown bool               `json:"pushedDown"` /* pushed down from higher query level? */
}
