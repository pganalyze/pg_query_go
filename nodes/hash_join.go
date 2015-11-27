// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

type HashJoin struct {
	Join        Join   `json:"join"`
	Hashclauses []Node `json:"hashclauses"`
}

func (node HashJoin) MarshalJSON() ([]byte, error) {
	type HashJoinMarshalAlias HashJoin
	return json.Marshal(map[string]interface{}{
		"HASHJOIN": (*HashJoinMarshalAlias)(&node),
	})
}

func (node *HashJoin) UnmarshalJSON(input []byte) (err error) {
	var fields map[string]json.RawMessage

	err = json.Unmarshal(input, &fields)
	if err != nil {
		return
	}

	if fields["join"] != nil {
		err = json.Unmarshal(fields["join"], &node.Join)
		if err != nil {
			return
		}
	}

	if fields["hashclauses"] != nil {
		node.Hashclauses, err = UnmarshalNodeArrayJSON(fields["hashclauses"])
		if err != nil {
			return
		}
	}

	return
}
