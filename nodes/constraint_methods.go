// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

func (node Constraint) MarshalJSON() ([]byte, error) {
	type ConstraintMarshalAlias Constraint
	return json.Marshal(map[string]interface{}{
		"CONSTRAINT": (*ConstraintMarshalAlias)(&node),
	})
}

func (node *Constraint) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}

func (node Constraint) Deparse() string {
	panic("Not Implemented")
}
