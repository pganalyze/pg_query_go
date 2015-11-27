// Auto-generated from postgres/src/include/nodes/primnodes.h - DO NOT EDIT

package pg_query

import "encoding/json"

/*
 * BooleanTest
 *
 * BooleanTest represents the operation of determining whether a boolean
 * is TRUE, FALSE, or UNKNOWN (ie, NULL).  All six meaningful combinations
 * are supported.  Note that a NULL input does *not* cause a NULL result.
 * The appropriate test is performed and returned as a boolean Datum.
 */
type BooleanTest struct {
	Xpr          Expr         `json:"xpr"`
	Arg          *Expr        `json:"arg"`          /* input expression */
	Booltesttype BoolTestType `json:"booltesttype"` /* test type */
}

func (node BooleanTest) MarshalJSON() ([]byte, error) {
	type BooleanTestMarshalAlias BooleanTest
	return json.Marshal(map[string]interface{}{
		"BOOLEANTEST": (*BooleanTestMarshalAlias)(&node),
	})
}

func (node *BooleanTest) UnmarshalJSON(input []byte) (err error) {
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

	if fields["arg"] != nil {
		var nodePtr *Node
		nodePtr, err = UnmarshalNodePtrJSON(fields["arg"])
		if err != nil {
			return
		}
		if nodePtr != nil && *nodePtr != nil {
			val := (*nodePtr).(Expr)
			node.Arg = &val
		}
	}

	if fields["booltesttype"] != nil {
		err = json.Unmarshal(fields["booltesttype"], &node.Booltesttype)
		if err != nil {
			return
		}
	}

	return
}
