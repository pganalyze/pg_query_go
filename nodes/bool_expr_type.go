// Auto-generated from postgres/src/include/nodes/primnodes.h - DO NOT EDIT

package pg_query

/*
 * BoolExpr - expression node for the basic Boolean operators AND, OR, NOT
 *
 * Notice the arguments are given as a List.  For NOT, of course the list
 * must always have exactly one element.  For AND and OR, the executor can
 * handle any number of arguments.  The parser generally treats AND and OR
 * as binary and so it typically only produces two-element lists, but the
 * optimizer will flatten trees of AND and OR nodes to produce longer lists
 * when possible.  There are also a few special cases where more arguments
 * can appear before optimization.
 */
type BoolExprType uint

const (
	AND_EXPR BoolExprType = iota
	OR_EXPR
)
