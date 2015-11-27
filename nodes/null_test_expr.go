// Auto-generated from postgres/src/include/nodes/primnodes.h - DO NOT EDIT

package pg_query

import "encoding/json"

/* ----------------
 * NullTest
 *
 * NullTest represents the operation of testing a value for NULLness.
 * The appropriate test is performed and returned as a boolean Datum.
 *
 * NOTE: the semantics of this for rowtype inputs are noticeably different
 * from the scalar case.  We provide an "argisrow" flag to reflect that.
 * ----------------
 */
type NullTest struct {
	Xpr          Expr         `json:"xpr"`
	Arg          *Expr        `json:"arg"`          /* input expression */
	Nulltesttype NullTestType `json:"nulltesttype"` /* IS NULL, IS NOT NULL */
	Argisrow     bool         `json:"argisrow"`     /* T if input is of a composite type */
}

func (node NullTest) MarshalJSON() ([]byte, error) {
	type NullTestMarshalAlias NullTest
	return json.Marshal(map[string]interface{}{
		"NULLTEST": (*NullTestMarshalAlias)(&node),
	})
}

func (node *NullTest) UnmarshalJSON(input []byte) (err error) {
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

	if fields["nulltesttype"] != nil {
		err = json.Unmarshal(fields["nulltesttype"], &node.Nulltesttype)
		if err != nil {
			return
		}
	}

	if fields["argisrow"] != nil {
		err = json.Unmarshal(fields["argisrow"], &node.Argisrow)
		if err != nil {
			return
		}
	}

	return
}
