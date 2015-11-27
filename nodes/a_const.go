package pg_query

import "encoding/json"

type A_Const struct {
	Type     string `json:"type"`
	Val      Value  `json:"val"`      /* value (includes type info, see value.h) */
	Location int    `json:"location"` /* token location, or -1 if unknown */
}

func (node A_Const) MarshalJSON() ([]byte, error) {
	if node.Type == "" {
		switch node.Val.Type {
		case T_Integer:
			node.Type = "integer"
		case T_Float:
			node.Type = "float"
		case T_String:
			node.Type = "string"
		case T_BitString:
			node.Type = "bit_string"
		case T_Null:
			node.Type = "null"
		}
	}

	type A_ConstMarshalAlias A_Const
	return json.Marshal(map[string]interface{}{
		"A_CONST": (*A_ConstMarshalAlias)(&node),
	})
}

func (node *A_Const) UnmarshalJSON(input []byte) (err error) {
	var fields map[string]json.RawMessage

	err = json.Unmarshal(input, &fields)
	if err != nil {
		return
	}

	if fields["type"] != nil {
		err = json.Unmarshal(fields["type"], &node.Type)
		if err != nil {
			return
		}
	}

	if fields["val"] != nil {
		err = json.Unmarshal(fields["val"], &node.Val)
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
