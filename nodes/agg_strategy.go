// Auto-generated - DO NOT EDIT

package pg_query

type AggStrategy uint

const (
	AGG_PLAIN  = iota /* simple agg across all input rows */
	AGG_SORTED        /* grouped agg, input must be sorted */
	AGG_HASHED        /* grouped agg, use internal hashtable */
)
