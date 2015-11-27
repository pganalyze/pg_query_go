// Auto-generated from postgres/src/include/nodes/parsenodes.h - DO NOT EDIT

package pg_query

/* ----------------------
 *		Vacuum and Analyze Statements
 *
 * Even though these are nominally two statements, it's convenient to use
 * just one node type for both.  Note that at least one of VACOPT_VACUUM
 * and VACOPT_ANALYZE must be set in options.  VACOPT_FREEZE is an internal
 * convenience for the grammar and is not examined at runtime --- the
 * freeze_min_age and freeze_table_age fields are what matter.
 * ----------------------
 */
type VacuumOption uint

const (
	VACOPT_VACUUM VacuumOption = iota
	VACOPT_ANALYZE
	VACOPT_VERBOSE
	VACOPT_FREEZE
	VACOPT_FULL
	VACOPT_NOWAIT
)
