// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

func (node SubLink) MarshalJSON() ([]byte, error) {
	type SubLinkMarshalAlias SubLink
	return json.Marshal(map[string]interface{}{
		"SUBLINK": (*SubLinkMarshalAlias)(&node),
	})
}

func (node *SubLink) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}

func (node SubLink) Deparse() string {
	panic("Not Implemented")
}
