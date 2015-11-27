// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

type PlaceHolderVar struct {
	Xpr        Expr     `json:"xpr"`
	Phexpr     *Expr    `json:"phexpr"`     /* the represented expression */
	Phrels     []uint32 `json:"phrels"`     /* base relids syntactically within expr src */
	Phid       Index    `json:"phid"`       /* ID for PHV (unique within planner run) */
	Phlevelsup Index    `json:"phlevelsup"` /* > 0 if PHV belongs to outer query */
}

func (node PlaceHolderVar) MarshalJSON() ([]byte, error) {
	type PlaceHolderVarMarshalAlias PlaceHolderVar
	return json.Marshal(map[string]interface{}{
		"PLACEHOLDERVAR": (*PlaceHolderVarMarshalAlias)(&node),
	})
}

func (node *PlaceHolderVar) UnmarshalJSON(input []byte) (err error) {
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

	if fields["phexpr"] != nil {
		var nodePtr *Node
		nodePtr, err = UnmarshalNodePtrJSON(fields["phexpr"])
		if err != nil {
			return
		}
		if nodePtr != nil && *nodePtr != nil {
			val := (*nodePtr).(Expr)
			node.Phexpr = &val
		}
	}

	if fields["phrels"] != nil {
		err = json.Unmarshal(fields["phrels"], &node.Phrels)
		if err != nil {
			return
		}
	}

	if fields["phid"] != nil {
		err = json.Unmarshal(fields["phid"], &node.Phid)
		if err != nil {
			return
		}
	}

	if fields["phlevelsup"] != nil {
		err = json.Unmarshal(fields["phlevelsup"], &node.Phlevelsup)
		if err != nil {
			return
		}
	}

	return
}
