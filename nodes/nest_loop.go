// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

type NestLoop struct {
	Join       Join   `json:"join"`
	NestParams []Node `json:"nestParams"` /* list of NestLoopParam nodes */
}

func (node NestLoop) MarshalJSON() ([]byte, error) {
	type NestLoopMarshalAlias NestLoop
	return json.Marshal(map[string]interface{}{
		"NESTLOOP": (*NestLoopMarshalAlias)(&node),
	})
}

func (node *NestLoop) UnmarshalJSON(input []byte) (err error) {
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

	if fields["nestParams"] != nil {
		node.NestParams, err = UnmarshalNodeArrayJSON(fields["nestParams"])
		if err != nil {
			return
		}
	}

	return
}
