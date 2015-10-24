// Auto-generated - DO NOT EDIT

package pg_query

type LockingClause struct {
	LockedRels []Node             `json:"lockedRels"` /* FOR [KEY] UPDATE/SHARE relations */
	Strength   LockClauseStrength `json:"strength"`
	NoWait     bool               `json:"noWait"` /* NOWAIT option */
}
