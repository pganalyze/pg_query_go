// Auto-generated from postgres/src/include/nodes/plannodes.h - DO NOT EDIT

package pg_query

type SetOpStrategy uint

const (
	SETOP_SORTED = iota /* input must be sorted */
	SETOP_HASHED        /* use internal hashtable */

)
