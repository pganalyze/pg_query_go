// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

type CollateClause struct {
	Arg      Node   `json:"arg"`      /* input expression */
	Collname []Node `json:"collname"` /* possibly-qualified collation name */
	Location int    `json:"location"` /* token location, or -1 if unknown */
}

func (node CollateClause) MarshalJSON() ([]byte, error) {
	type CollateClauseMarshalAlias CollateClause
	return json.Marshal(map[string]interface{}{
		"COLLATECLAUSE": (*CollateClauseMarshalAlias)(&node),
	})
}

func (node *CollateClause) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}
