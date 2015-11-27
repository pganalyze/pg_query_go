// Auto-generated from postgres/src/include/nodes/plannodes.h - DO NOT EDIT

package pg_query

/*
 * RowMarkType -
 *	  enums for types of row-marking operations
 *
 * The first four of these values represent different lock strengths that
 * we can take on tuples according to SELECT FOR [KEY] UPDATE/SHARE requests.
 * We only support these on regular tables.  For foreign tables, any locking
 * that might be done for these requests must happen during the initial row
 * fetch; there is no mechanism for going back to lock a row later (and thus
 * no need for EvalPlanQual machinery during updates of foreign tables).
 * This means that the semantics will be a bit different than for a local
 * table; in particular we are likely to lock more rows than would be locked
 * locally, since remote rows will be locked even if they then fail
 * locally-checked restriction or join quals.  However, the alternative of
 * doing a separate remote query to lock each selected row is extremely
 * unappealing, so let's do it like this for now.
 *
 * When doing UPDATE, DELETE, or SELECT FOR UPDATE/SHARE, we have to uniquely
 * identify all the source rows, not only those from the target relations, so
 * that we can perform EvalPlanQual rechecking at need.  For plain tables we
 * can just fetch the TID, much as for a target relation; this case is
 * represented by ROW_MARK_REFERENCE.  Otherwise (for example for VALUES or
 * FUNCTION scans) we have to copy the whole row value.  ROW_MARK_COPY is
 * pretty inefficient, since most of the time we'll never need the data; but
 * fortunately the case is not performance-critical in practice.  Note that
 * we use ROW_MARK_COPY for non-target foreign tables, even if the FDW has a
 * concept of rowid and so could theoretically support some form of
 * ROW_MARK_REFERENCE.  Although copying the whole row value is inefficient,
 * it's probably still faster than doing a second remote fetch, so it doesn't
 * seem worth the extra complexity to permit ROW_MARK_REFERENCE.
 */
type RowMarkType uint

const (
	ROW_MARK_EXCLUSIVE      RowMarkType = iota /* obtain exclusive tuple lock */
	ROW_MARK_NOKEYEXCLUSIVE                    /* obtain no-key exclusive tuple lock */
	ROW_MARK_SHARE                             /* obtain shared tuple lock */
	ROW_MARK_KEYSHARE                          /* obtain keyshare tuple lock */
	ROW_MARK_REFERENCE                         /* just fetch the TID */
	ROW_MARK_COPY                              /* physically copy the row value */
)
