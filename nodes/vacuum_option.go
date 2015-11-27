// Auto-generated from postgres/src/include/nodes/parsenodes.h - DO NOT EDIT

package pg_query

type VacuumOption uint

const (
	VACOPT_VACUUM = iota
	VACOPT_ANALYZE
	VACOPT_VERBOSE
	VACOPT_FREEZE
	VACOPT_FULL
	VACOPT_NOWAIT
)
