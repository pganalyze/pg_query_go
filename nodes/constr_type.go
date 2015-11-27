// Auto-generated from postgres/src/include/nodes/parsenodes.h - DO NOT EDIT

package pg_query

type ConstrType uint

const (
	CONSTR_NULL = iota /* not standard SQL, but a lot of people
	 * expect it */

	CONSTR_NOTNULL
	CONSTR_DEFAULT
	CONSTR_CHECK
	CONSTR_PRIMARY
	CONSTR_UNIQUE
	CONSTR_EXCLUSION
	CONSTR_FOREIGN
	CONSTR_ATTR_DEFERRABLE /* attributes for previous constraint node */
	CONSTR_ATTR_NOT_DEFERRABLE
	CONSTR_ATTR_DEFERRED
	CONSTR_ATTR_IMMEDIATE
)
