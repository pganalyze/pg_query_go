// Auto-generated from postgres/src/include/nodes/primnodes.h - DO NOT EDIT

package pg_query

type InhOption uint

const (
	INH_NO      InhOption = iota /* Do NOT scan child tables */
	INH_YES                      /* DO scan child tables */
	INH_DEFAULT                  /* Use current SQL_inheritance option */
)
