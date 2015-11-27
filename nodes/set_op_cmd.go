// Auto-generated from postgres/src/include/nodes/plannodes.h - DO NOT EDIT

package pg_query

/* ----------------
 *		setop node
 * ----------------
 */
type SetOpCmd uint

const (
	SETOPCMD_INTERSECT SetOpCmd = iota
	SETOPCMD_INTERSECT_ALL
	SETOPCMD_EXCEPT
	SETOPCMD_EXCEPT_ALL
)
