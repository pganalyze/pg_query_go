// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

type PlaceHolderVar struct {
	Xpr        Expr     `json:"xpr"`
	Phexpr     *Expr    `json:"phexpr"`     /* the represented expression */
	Phrels     []uint32 `json:"phrels"`     /* base relids syntactically within expr src */
	Phid       Index    `json:"phid"`       /* ID for PHV (unique within planner run) */
	Phlevelsup Index    `json:"phlevelsup"` /* > 0 if PHV belongs to outer query */
}

func (node PlaceHolderVar) MarshalJSON() ([]byte, error) {
	type PlaceHolderVarMarshalAlias PlaceHolderVar
	return json.Marshal(map[string]interface{}{
		"PLACEHOLDERVAR": (*PlaceHolderVarMarshalAlias)(&node),
	})
}

func (node *PlaceHolderVar) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}
