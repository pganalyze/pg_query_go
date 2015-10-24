// Auto-generated - DO NOT EDIT

package pg_query

type CoercionContext uint

const (
	COERCION_IMPLICIT   = iota /* coercion in context of expression */
	COERCION_ASSIGNMENT        /* coercion in context of assignment */
	COERCION_EXPLICIT          /* explicit cast operation */
)
