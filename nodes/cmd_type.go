// Auto-generated - DO NOT EDIT

package pg_query

type CmdType uint

const (
	CMD_UNKNOWN = iota
	CMD_SELECT  /* select stmt */
	CMD_UPDATE  /* update stmt */
	CMD_INSERT  /* insert stmt */
	CMD_DELETE
	CMD_UTILITY /* cmds like create, destroy, copy, vacuum,
	 * etc. */
	CMD_NOTHING /* dummy command for instead nothing rules
	 * with qual */

)
