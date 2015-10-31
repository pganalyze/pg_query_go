// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

func (node Query) MarshalJSON() ([]byte, error) {
	type QueryMarshalAlias Query
	return json.Marshal(map[string]interface{}{
		"QUERY": (*QueryMarshalAlias)(&node),
	})
}

func (node *Query) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}

func (node Query) Deparse() string {
	panic("Not Implemented")
}
