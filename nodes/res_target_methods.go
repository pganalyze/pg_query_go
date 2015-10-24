package pg_query

import "encoding/json"

func (input ResTarget) MarshalJSON() ([]byte, error) {
	type ResTargetAlias ResTarget
	return json.Marshal(map[string]interface{}{
		"RESTARGET": (*ResTargetAlias)(&input),
	})
}

func (resTarget *ResTarget) UnmarshalJSON(input []byte) error {
	return UnmarshalNodeFieldJSON(input, resTarget)
}

func (resTarget ResTarget) Deparse() string {
	panic("Not Implemented")
}
