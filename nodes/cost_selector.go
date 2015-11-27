// Auto-generated from postgres/src/include/nodes/relation.h - DO NOT EDIT

package pg_query

/*
 * When looking for a "cheapest path", this enum specifies whether we want
 * cheapest startup cost or cheapest total cost.
 */
type CostSelector uint

const (
	STARTUP_COST CostSelector = iota
	TOTAL_COST
)
