package pg_query

type InhOption uint

const (
  INH_NO = iota						/* Do NOT scan child tables */
  INH_YES					/* DO scan child tables */
  INH_DEFAULT					/* Use current SQL_inheritance option */
)
