// Auto-generated from postgres/src/include/nodes/parsenodes.h - DO NOT EDIT

package pg_query

type LockClauseStrength uint

const (
	/* order is important -- see applyLockingClause */
	LCS_FORKEYSHARE = iota
	LCS_FORSHARE
	LCS_FORNOKEYUPDATE
	LCS_FORUPDATE
)
