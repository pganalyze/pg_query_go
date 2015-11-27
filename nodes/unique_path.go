// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

type UniquePath struct {
	Path        Path             `json:"path"`
	Subpath     *Path            `json:"subpath"`
	Umethod     UniquePathMethod `json:"umethod"`
	InOperators []Node           `json:"in_operators"` /* equality operators of the IN clause */
	UniqExprs   []Node           `json:"uniq_exprs"`   /* expressions to be made unique */
}

func (node UniquePath) MarshalJSON() ([]byte, error) {
	type UniquePathMarshalAlias UniquePath
	return json.Marshal(map[string]interface{}{
		"UNIQUEPATH": (*UniquePathMarshalAlias)(&node),
	})
}

func (node *UniquePath) UnmarshalJSON(input []byte) (err error) {
	var fields map[string]json.RawMessage

	err = json.Unmarshal(input, &fields)
	if err != nil {
		return
	}

	if fields["path"] != nil {
		err = json.Unmarshal(fields["path"], &node.Path)
		if err != nil {
			return
		}
	}

	if fields["subpath"] != nil {
		var nodePtr *Node
		nodePtr, err = UnmarshalNodePtrJSON(fields["subpath"])
		if err != nil {
			return
		}
		if nodePtr != nil && *nodePtr != nil {
			val := (*nodePtr).(Path)
			node.Subpath = &val
		}
	}

	if fields["umethod"] != nil {
		err = json.Unmarshal(fields["umethod"], &node.Umethod)
		if err != nil {
			return
		}
	}

	if fields["in_operators"] != nil {
		node.InOperators, err = UnmarshalNodeArrayJSON(fields["in_operators"])
		if err != nil {
			return
		}
	}

	if fields["uniq_exprs"] != nil {
		node.UniqExprs, err = UnmarshalNodeArrayJSON(fields["uniq_exprs"])
		if err != nil {
			return
		}
	}

	return
}
