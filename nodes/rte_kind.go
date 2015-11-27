// Auto-generated from postgres/src/include/nodes/parsenodes.h - DO NOT EDIT

package pg_query

type RTEKind uint

const (
	RTE_RELATION = iota /* ordinary relation reference */
	RTE_SUBQUERY        /* subquery in FROM */
	RTE_JOIN            /* join */
	RTE_FUNCTION        /* function in FROM */
	RTE_VALUES          /* VALUES (<exprlist>), (<exprlist>), ... */
	RTE_CTE             /* common table expr (WITH list element) */

)
