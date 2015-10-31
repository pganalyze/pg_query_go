// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

func (node EquivalenceClass) MarshalJSON() ([]byte, error) {
	type EquivalenceClassMarshalAlias EquivalenceClass
	return json.Marshal(map[string]interface{}{
		"EQUIVALENCECLASS": (*EquivalenceClassMarshalAlias)(&node),
	})
}

func (node *EquivalenceClass) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}

func (node EquivalenceClass) Deparse() string {
	panic("Not Implemented")
}
