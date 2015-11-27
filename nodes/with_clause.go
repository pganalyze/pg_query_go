// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

type WithClause struct {
	Ctes      []Node `json:"ctes"`      /* list of CommonTableExprs */
	Recursive bool   `json:"recursive"` /* true = WITH RECURSIVE */
	Location  int    `json:"location"`  /* token location, or -1 if unknown */
}

func (node WithClause) MarshalJSON() ([]byte, error) {
	type WithClauseMarshalAlias WithClause
	return json.Marshal(map[string]interface{}{
		"WITHCLAUSE": (*WithClauseMarshalAlias)(&node),
	})
}

func (node *WithClause) UnmarshalJSON(input []byte) (err error) {
	var fields map[string]json.RawMessage

	err = json.Unmarshal(input, &fields)
	if err != nil {
		return
	}

	if fields["ctes"] != nil {
		node.Ctes, err = UnmarshalNodeArrayJSON(fields["ctes"])
		if err != nil {
			return
		}
	}

	if fields["recursive"] != nil {
		err = json.Unmarshal(fields["recursive"], &node.Recursive)
		if err != nil {
			return
		}
	}

	if fields["location"] != nil {
		err = json.Unmarshal(fields["location"], &node.Location)
		if err != nil {
			return
		}
	}

	return
}
