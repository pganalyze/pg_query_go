// Auto-generated from postgres/src/include/nodes/primnodes.h - DO NOT EDIT

package pg_query

import "encoding/json"

type CaseWhen struct {
	Xpr      Expr  `json:"xpr"`
	Expr     *Expr `json:"expr"`     /* condition expression */
	Result   *Expr `json:"result"`   /* substitution result */
	Location int   `json:"location"` /* token location, or -1 if unknown */
}

func (node CaseWhen) MarshalJSON() ([]byte, error) {
	type CaseWhenMarshalAlias CaseWhen
	return json.Marshal(map[string]interface{}{
		"WHEN": (*CaseWhenMarshalAlias)(&node),
	})
}

func (node *CaseWhen) UnmarshalJSON(input []byte) (err error) {
	var fields map[string]json.RawMessage

	err = json.Unmarshal(input, &fields)
	if err != nil {
		return
	}

	if fields["xpr"] != nil {
		err = json.Unmarshal(fields["xpr"], &node.Xpr)
		if err != nil {
			return
		}
	}

	if fields["expr"] != nil {
		var nodePtr *Node
		nodePtr, err = UnmarshalNodePtrJSON(fields["expr"])
		if err != nil {
			return
		}
		if nodePtr != nil && *nodePtr != nil {
			val := (*nodePtr).(Expr)
			node.Expr = &val
		}
	}

	if fields["result"] != nil {
		var nodePtr *Node
		nodePtr, err = UnmarshalNodePtrJSON(fields["result"])
		if err != nil {
			return
		}
		if nodePtr != nil && *nodePtr != nil {
			val := (*nodePtr).(Expr)
			node.Result = &val
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
