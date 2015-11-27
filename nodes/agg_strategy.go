// Auto-generated from postgres/src/include/nodes/plannodes.h - DO NOT EDIT

package pg_query

/* ---------------
 *		aggregate node
 *
 * An Agg node implements plain or grouped aggregation.  For grouped
 * aggregation, we can work with presorted input or unsorted input;
 * the latter strategy uses an internal hashtable.
 *
 * Notice the lack of any direct info about the aggregate functions to be
 * computed.  They are found by scanning the node's tlist and quals during
 * executor startup.  (It is possible that there are no aggregate functions;
 * this could happen if they get optimized away by constant-folding, or if
 * we are using the Agg node to implement hash-based grouping.)
 * ---------------
 */
type AggStrategy uint

const (
	AGG_PLAIN  AggStrategy = iota /* simple agg across all input rows */
	AGG_SORTED                    /* grouped agg, input must be sorted */
	AGG_HASHED                    /* grouped agg, use internal hashtable */
)
