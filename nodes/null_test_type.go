// Auto-generated from postgres/src/include/nodes/primnodes.h - DO NOT EDIT

package pg_query

/* ----------------
 * NullTest
 *
 * NullTest represents the operation of testing a value for NULLness.
 * The appropriate test is performed and returned as a boolean Datum.
 *
 * NOTE: the semantics of this for rowtype inputs are noticeably different
 * from the scalar case.  We provide an "argisrow" flag to reflect that.
 * ----------------
 */
type NullTestType uint

const (
	IS_NULL NullTestType = iota
	IS_NOT_NULL
)
