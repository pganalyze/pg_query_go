// Auto-generated - DO NOT EDIT

package pg_query

type PlaceHolderVar struct {
	Xpr        Expr     `json:"xpr"`
	Phexpr     *Expr    `json:"phexpr"`     /* the represented expression */
	Phrels     []uint32 `json:"phrels"`     /* base relids syntactically within expr src */
	Phid       Index    `json:"phid"`       /* ID for PHV (unique within planner run) */
	Phlevelsup Index    `json:"phlevelsup"` /* > 0 if PHV belongs to outer query */
}
