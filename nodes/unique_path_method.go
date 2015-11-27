// Auto-generated from postgres/src/include/nodes/relation.h - DO NOT EDIT

package pg_query

/*
 * UniquePath represents elimination of distinct rows from the output of
 * its subpath.
 *
 * This is unlike the other Path nodes in that it can actually generate
 * different plans: either hash-based or sort-based implementation, or a
 * no-op if the input path can be proven distinct already.  The decision
 * is sufficiently localized that it's not worth having separate Path node
 * types.  (Note: in the no-op case, we could eliminate the UniquePath node
 * entirely and just return the subpath; but it's convenient to have a
 * UniquePath in the path tree to signal upper-level routines that the input
 * is known distinct.)
 */
type UniquePathMethod uint

const (
	UNIQUE_PATH_NOOP UniquePathMethod = iota /* input is known unique already */
	UNIQUE_PATH_HASH                         /* use hashing */
	UNIQUE_PATH_SORT                         /* use sorting */
)
