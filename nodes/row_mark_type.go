// Auto-generated - DO NOT EDIT

package pg_query

type RowMarkType uint

const (
	ROW_MARK_EXCLUSIVE      = iota /* obtain exclusive tuple lock */
	ROW_MARK_NOKEYEXCLUSIVE        /* obtain no-key exclusive tuple lock */
	ROW_MARK_SHARE                 /* obtain shared tuple lock */
	ROW_MARK_KEYSHARE              /* obtain keyshare tuple lock */
	ROW_MARK_REFERENCE             /* just fetch the TID */
	ROW_MARK_COPY                  /* physically copy the row value */
)
