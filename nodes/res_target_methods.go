// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

func (node ResTarget) MarshalJSON() ([]byte, error) {
	type ResTargetMarshalAlias ResTarget
	return json.Marshal(map[string]interface{}{
		"RESTARGET": (*ResTargetMarshalAlias)(&node),
	})
}

func (node *ResTarget) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}

func (node ResTarget) Deparse() string {
	panic("Not Implemented")
}
