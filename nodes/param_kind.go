// Auto-generated from postgres/src/include/nodes/primnodes.h - DO NOT EDIT

package pg_query

/* ----------------
 * Param
 *		paramkind - specifies the kind of parameter. The possible values
 *		for this field are:
 *
 *		PARAM_EXTERN:  The parameter value is supplied from outside the plan.
 *				Such parameters are numbered from 1 to n.
 *
 *		PARAM_EXEC:  The parameter is an internal executor parameter, used
 *				for passing values into and out of sub-queries or from
 *				nestloop joins to their inner scans.
 *				For historical reasons, such parameters are numbered from 0.
 *				These numbers are independent of PARAM_EXTERN numbers.
 *
 *		PARAM_SUBLINK:	The parameter represents an output column of a SubLink
 *				node's sub-select.  The column number is contained in the
 *				`paramid' field.  (This type of Param is converted to
 *				PARAM_EXEC during planning.)
 *
 * Note: currently, paramtypmod is valid for PARAM_SUBLINK Params, and for
 * PARAM_EXEC Params generated from them; it is always -1 for PARAM_EXTERN
 * params, since the APIs that supply values for such parameters don't carry
 * any typmod info.
 * ----------------
 */
type ParamKind uint

const (
	PARAM_EXTERN ParamKind = iota
	PARAM_EXEC
	PARAM_SUBLINK
)
