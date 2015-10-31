// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

func (node Alias) MarshalJSON() ([]byte, error) {
	type AliasMarshalAlias Alias
	return json.Marshal(map[string]interface{}{
		"ALIAS": (*AliasMarshalAlias)(&node),
	})
}

func (node *Alias) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}

func (node Alias) Deparse() string {
	panic("Not Implemented")
}
