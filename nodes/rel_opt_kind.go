// Auto-generated from postgres/src/include/nodes/relation.h - DO NOT EDIT

package pg_query

type RelOptKind uint

const (
	RELOPT_BASEREL = iota
	RELOPT_JOINREL
	RELOPT_OTHER_MEMBER_REL
	RELOPT_DEADREL
)
