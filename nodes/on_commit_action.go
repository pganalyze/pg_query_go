// Auto-generated from postgres/src/include/nodes/primnodes.h - DO NOT EDIT

package pg_query

type OnCommitAction uint

const (
	ONCOMMIT_NOOP          = iota /* No ON COMMIT clause (do nothing) */
	ONCOMMIT_PRESERVE_ROWS        /* ON COMMIT PRESERVE ROWS (do nothing) */
	ONCOMMIT_DELETE_ROWS          /* ON COMMIT DELETE ROWS */
	ONCOMMIT_DROP                 /* ON COMMIT DROP */

)
