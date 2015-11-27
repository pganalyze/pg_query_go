// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

type Limit struct {
	Plan        Plan `json:"plan"`
	LimitOffset Node `json:"limitOffset"` /* OFFSET parameter, or NULL if none */
	LimitCount  Node `json:"limitCount"`  /* COUNT parameter, or NULL if none */
}

func (node Limit) MarshalJSON() ([]byte, error) {
	type LimitMarshalAlias Limit
	return json.Marshal(map[string]interface{}{
		"LIMIT": (*LimitMarshalAlias)(&node),
	})
}

func (node *Limit) UnmarshalJSON(input []byte) (err error) {
	var fields map[string]json.RawMessage

	err = json.Unmarshal(input, &fields)
	if err != nil {
		return
	}

	if fields["plan"] != nil {
		err = json.Unmarshal(fields["plan"], &node.Plan)
		if err != nil {
			return
		}
	}

	if fields["limitOffset"] != nil {
		node.LimitOffset, err = UnmarshalNodeJSON(fields["limitOffset"])
		if err != nil {
			return
		}
	}

	if fields["limitCount"] != nil {
		node.LimitCount, err = UnmarshalNodeJSON(fields["limitCount"])
		if err != nil {
			return
		}
	}

	return
}
