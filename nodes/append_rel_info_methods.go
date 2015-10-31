// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

func (node AppendRelInfo) MarshalJSON() ([]byte, error) {
	type AppendRelInfoMarshalAlias AppendRelInfo
	return json.Marshal(map[string]interface{}{
		"APPENDRELINFO": (*AppendRelInfoMarshalAlias)(&node),
	})
}

func (node *AppendRelInfo) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}

func (node AppendRelInfo) Deparse() string {
	panic("Not Implemented")
}
