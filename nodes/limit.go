// Auto-generated - DO NOT EDIT

package pg_query

type Limit struct {
	Plan        Plan `json:"plan"`
	LimitOffset Node `json:"limitOffset"` /* OFFSET parameter, or NULL if none */
	LimitCount  Node `json:"limitCount"`  /* COUNT parameter, or NULL if none */
}
