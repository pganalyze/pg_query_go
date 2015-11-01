// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

type Alias struct {
	Aliasname *string `json:"aliasname"` /* aliased rel name (never qualified) */
	Colnames  []Node  `json:"colnames"`  /* optional list of column aliases */
}

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
