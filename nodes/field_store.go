// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

type FieldStore struct {
	Xpr        Expr   `json:"xpr"`
	Arg        *Expr  `json:"arg"`        /* input tuple value */
	Newvals    []Node `json:"newvals"`    /* new value(s) for field(s) */
	Fieldnums  []Node `json:"fieldnums"`  /* integer list of field attnums */
	Resulttype Oid    `json:"resulttype"` /* type of result (same as type of arg) */
	/* Like RowExpr, we deliberately omit a typmod and collation here */
}

func (node FieldStore) MarshalJSON() ([]byte, error) {
	type FieldStoreMarshalAlias FieldStore
	return json.Marshal(map[string]interface{}{
		"FIELDSTORE": (*FieldStoreMarshalAlias)(&node),
	})
}

func (node *FieldStore) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}
