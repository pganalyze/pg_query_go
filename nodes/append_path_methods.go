// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

func (node AppendPath) MarshalJSON() ([]byte, error) {
	type AppendPathMarshalAlias AppendPath
	return json.Marshal(map[string]interface{}{
		"APPENDPATH": (*AppendPathMarshalAlias)(&node),
	})
}

func (node *AppendPath) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}

func (node AppendPath) Deparse() string {
	panic("Not Implemented")
}
