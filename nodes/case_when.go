// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

type CaseWhen struct {
	Xpr      Expr  `json:"xpr"`
	Expr     *Expr `json:"expr"`     /* condition expression */
	Result   *Expr `json:"result"`   /* substitution result */
	Location int   `json:"location"` /* token location, or -1 if unknown */
}

func (node CaseWhen) MarshalJSON() ([]byte, error) {
	type CaseWhenMarshalAlias CaseWhen
	return json.Marshal(map[string]interface{}{
		"WHEN": (*CaseWhenMarshalAlias)(&node),
	})
}

func (node *CaseWhen) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}
