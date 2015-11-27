// Auto-generated from postgres/src/include/nodes/primnodes.h - DO NOT EDIT

package pg_query

import "encoding/json"

/*
 * NamedArgExpr - a named argument of a function
 *
 * This node type can only appear in the args list of a FuncCall or FuncExpr
 * node.  We support pure positional call notation (no named arguments),
 * named notation (all arguments are named), and mixed notation (unnamed
 * arguments followed by named ones).
 *
 * Parse analysis sets argnumber to the positional index of the argument,
 * but doesn't rearrange the argument list.
 *
 * The planner will convert argument lists to pure positional notation
 * during expression preprocessing, so execution never sees a NamedArgExpr.
 */
type NamedArgExpr struct {
	Xpr       Expr    `json:"xpr"`
	Arg       *Expr   `json:"arg"`       /* the argument expression */
	Name      *string `json:"name"`      /* the name */
	Argnumber int     `json:"argnumber"` /* argument's number in positional notation */
	Location  int     `json:"location"`  /* argument name location, or -1 if unknown */
}

func (node NamedArgExpr) MarshalJSON() ([]byte, error) {
	type NamedArgExprMarshalAlias NamedArgExpr
	return json.Marshal(map[string]interface{}{
		"NAMEDARGEXPR": (*NamedArgExprMarshalAlias)(&node),
	})
}

func (node *NamedArgExpr) UnmarshalJSON(input []byte) (err error) {
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

	if fields["name"] != nil {
		err = json.Unmarshal(fields["name"], &node.Name)
		if err != nil {
			return
		}
	}

	if fields["argnumber"] != nil {
		err = json.Unmarshal(fields["argnumber"], &node.Argnumber)
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
