// Auto-generated - DO NOT EDIT

package pg_query

type LockClauseStrength uint

const (
	/* order is important -- see applyLockingClause */
	LCS_FORKEYSHARE = iota
	LCS_FORSHARE
	LCS_FORNOKEYUPDATE
	LCS_FORUPDATE
)
