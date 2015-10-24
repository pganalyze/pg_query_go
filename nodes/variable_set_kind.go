// Auto-generated - DO NOT EDIT

package pg_query

type VariableSetKind uint

const (
	VAR_SET_VALUE   = iota /* SET var = value */
	VAR_SET_DEFAULT        /* SET var TO DEFAULT */
	VAR_SET_CURRENT        /* SET var FROM CURRENT */
	VAR_SET_MULTI          /* special case for SET TRANSACTION ... */
	VAR_RESET              /* RESET var */
	VAR_RESET_ALL          /* RESET ALL */
)
