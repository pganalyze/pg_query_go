// Auto-generated from postgres/src/include/nodes/primnodes.h - DO NOT EDIT

package pg_query

type CoercionForm uint

const (
	COERCE_EXPLICIT_CALL = iota /* display as a function call */
	COERCE_EXPLICIT_CAST        /* display as an explicit cast */
	COERCE_IMPLICIT_CAST        /* implicit cast, so hide it */

)
