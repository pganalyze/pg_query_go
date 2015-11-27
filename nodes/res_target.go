// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

type ResTarget struct {
	Name        *string `json:"name"`        /* column name or NULL */
	Indirection []Node  `json:"indirection"` /* subscripts, field names, and '*', or NIL */
	Val         Node    `json:"val"`         /* the value expression to compute or assign */
	Location    int     `json:"location"`    /* token location, or -1 if unknown */
}

func (node ResTarget) MarshalJSON() ([]byte, error) {
	type ResTargetMarshalAlias ResTarget
	return json.Marshal(map[string]interface{}{
		"RESTARGET": (*ResTargetMarshalAlias)(&node),
	})
}

func (node *ResTarget) UnmarshalJSON(input []byte) (err error) {
	var fields map[string]json.RawMessage

	err = json.Unmarshal(input, &fields)
	if err != nil {
		return
	}

	if fields["name"] != nil {
		err = json.Unmarshal(fields["name"], &node.Name)
		if err != nil {
			return
		}
	}

	if fields["indirection"] != nil {
		node.Indirection, err = UnmarshalNodeArrayJSON(fields["indirection"])
		if err != nil {
			return
		}
	}

	if fields["val"] != nil {
		node.Val, err = UnmarshalNodeJSON(fields["val"])
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
