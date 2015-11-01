// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

type UniquePath struct {
	Path        Path             `json:"path"`
	Subpath     *Path            `json:"subpath"`
	Umethod     UniquePathMethod `json:"umethod"`
	InOperators []Node           `json:"in_operators"` /* equality operators of the IN clause */
	UniqExprs   []Node           `json:"uniq_exprs"`   /* expressions to be made unique */
}

func (node UniquePath) MarshalJSON() ([]byte, error) {
	type UniquePathMarshalAlias UniquePath
	return json.Marshal(map[string]interface{}{
		"UNIQUEPATH": (*UniquePathMarshalAlias)(&node),
	})
}

func (node *UniquePath) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}
