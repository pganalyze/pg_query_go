// Auto-generated - DO NOT EDIT

package pg_query

type UniquePathMethod uint

const (
	UNIQUE_PATH_NOOP = iota /* input is known unique already */
	UNIQUE_PATH_HASH        /* use hashing */
	UNIQUE_PATH_SORT        /* use sorting */

)
