// Auto-generated - DO NOT EDIT

package pg_query

type A_Expr_Kind uint

const (
	AEXPR_OP  = iota /* normal operator */
	AEXPR_AND        /* booleans - name field is unused */
	AEXPR_OR
	AEXPR_NOT
	AEXPR_OP_ANY   /* scalar op ANY (array) */
	AEXPR_OP_ALL   /* scalar op ALL (array) */
	AEXPR_DISTINCT /* IS DISTINCT FROM - name must be "=" */
	AEXPR_NULLIF   /* NULLIF - name must be "=" */
	AEXPR_OF       /* IS [NOT] OF - name must be "=" or "<>" */
	AEXPR_IN       /* [NOT] IN - name must be "=" or "<>" */

)
