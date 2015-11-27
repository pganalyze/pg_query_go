// Auto-generated from postgres/src/include/nodes/parsenodes.h - DO NOT EDIT

package pg_query

type DiscardMode uint

const (
	DISCARD_ALL = iota
	DISCARD_PLANS
	DISCARD_SEQUENCES
	DISCARD_TEMP
)
