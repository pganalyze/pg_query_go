// Auto-generated - DO NOT EDIT

package pg_query

type LateralJoinInfo struct {
	LateralLhs []uint32 `json:"lateral_lhs"` /* rels needed to compute a lateral value */
	LateralRhs []uint32 `json:"lateral_rhs"` /* rel where lateral value is needed */
}
