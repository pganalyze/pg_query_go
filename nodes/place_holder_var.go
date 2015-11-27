// Auto-generated from postgres/src/include/nodes/relation.h - DO NOT EDIT

package pg_query

import "encoding/json"

/*
 * Placeholder node for an expression to be evaluated below the top level
 * of a plan tree.  This is used during planning to represent the contained
 * expression.  At the end of the planning process it is replaced by either
 * the contained expression or a Var referring to a lower-level evaluation of
 * the contained expression.  Typically the evaluation occurs below an outer
 * join, and Var references above the outer join might thereby yield NULL
 * instead of the expression value.
 *
 * Although the planner treats this as an expression node type, it is not
 * recognized by the parser or executor, so we declare it here rather than
 * in primnodes.h.
 */
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
