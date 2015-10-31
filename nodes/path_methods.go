// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

func (node Path) MarshalJSON() ([]byte, error) {
	type PathMarshalAlias Path
	return json.Marshal(map[string]interface{}{
		"PATH": (*PathMarshalAlias)(&node),
	})
}

func (node *Path) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}

func (node Path) Deparse() string {
	panic("Not Implemented")
}
