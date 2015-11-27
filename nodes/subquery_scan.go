// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

type SubqueryScan struct {
	Scan    Scan  `json:"scan"`
	Subplan *Plan `json:"subplan"`
}

func (node SubqueryScan) MarshalJSON() ([]byte, error) {
	type SubqueryScanMarshalAlias SubqueryScan
	return json.Marshal(map[string]interface{}{
		"SUBQUERYSCAN": (*SubqueryScanMarshalAlias)(&node),
	})
}

func (node *SubqueryScan) UnmarshalJSON(input []byte) (err error) {
	var fields map[string]json.RawMessage

	err = json.Unmarshal(input, &fields)
	if err != nil {
		return
	}

	if fields["scan"] != nil {
		err = json.Unmarshal(fields["scan"], &node.Scan)
		if err != nil {
			return
		}
	}

	if fields["subplan"] != nil {
		var nodePtr *Node
		nodePtr, err = UnmarshalNodePtrJSON(fields["subplan"])
		if err != nil {
			return
		}
		if nodePtr != nil && *nodePtr != nil {
			val := (*nodePtr).(Plan)
			node.Subplan = &val
		}
	}

	return
}
