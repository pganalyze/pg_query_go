// Auto-generated from postgres/src/include/nodes/parsenodes.h - DO NOT EDIT

package pg_query

type SortByDir uint

const (
	SORTBY_DEFAULT = iota
	SORTBY_ASC
	SORTBY_DESC
	SORTBY_USING /* not allowed in CREATE INDEX ... */
)
