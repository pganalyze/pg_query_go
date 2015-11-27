// Auto-generated from postgres/src/include/nodes/parsenodes.h - DO NOT EDIT

package pg_query

type SetOperation uint

const (
	SETOP_NONE = iota
	SETOP_UNION
	SETOP_INTERSECT
	SETOP_EXCEPT
)
