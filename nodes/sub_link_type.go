// Auto-generated from postgres/src/include/nodes/primnodes.h - DO NOT EDIT

package pg_query

type SubLinkType uint

const (
	EXISTS_SUBLINK = iota
	ALL_SUBLINK
	ANY_SUBLINK
	ROWCOMPARE_SUBLINK
	EXPR_SUBLINK
	ARRAY_SUBLINK
	CTE_SUBLINK /* for SubPlans only */
)
