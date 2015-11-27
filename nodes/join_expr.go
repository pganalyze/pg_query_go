// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

type JoinExpr struct {
	Jointype    JoinType `json:"jointype"`    /* type of join */
	IsNatural   bool     `json:"isNatural"`   /* Natural join? Will need to shape table */
	Larg        Node     `json:"larg"`        /* left subtree */
	Rarg        Node     `json:"rarg"`        /* right subtree */
	UsingClause []Node   `json:"usingClause"` /* USING clause, if any (list of String) */
	Quals       Node     `json:"quals"`       /* qualifiers on join, if any */
	Alias       *Alias   `json:"alias"`       /* user-written alias clause, if any */
	Rtindex     int      `json:"rtindex"`     /* RT index assigned for join, or 0 */
}

func (node JoinExpr) MarshalJSON() ([]byte, error) {
	type JoinExprMarshalAlias JoinExpr
	return json.Marshal(map[string]interface{}{
		"JOINEXPR": (*JoinExprMarshalAlias)(&node),
	})
}

func (node *JoinExpr) UnmarshalJSON(input []byte) (err error) {
	var fields map[string]json.RawMessage

	err = json.Unmarshal(input, &fields)
	if err != nil {
		return
	}

	if fields["jointype"] != nil {
		err = json.Unmarshal(fields["jointype"], &node.Jointype)
		if err != nil {
			return
		}
	}

	if fields["isNatural"] != nil {
		err = json.Unmarshal(fields["isNatural"], &node.IsNatural)
		if err != nil {
			return
		}
	}

	if fields["larg"] != nil {
		node.Larg, err = UnmarshalNodeJSON(fields["larg"])
		if err != nil {
			return
		}
	}

	if fields["rarg"] != nil {
		node.Rarg, err = UnmarshalNodeJSON(fields["rarg"])
		if err != nil {
			return
		}
	}

	if fields["usingClause"] != nil {
		node.UsingClause, err = UnmarshalNodeArrayJSON(fields["usingClause"])
		if err != nil {
			return
		}
	}

	if fields["quals"] != nil {
		node.Quals, err = UnmarshalNodeJSON(fields["quals"])
		if err != nil {
			return
		}
	}

	if fields["alias"] != nil {
		var nodePtr *Node
		nodePtr, err = UnmarshalNodePtrJSON(fields["alias"])
		if err != nil {
			return
		}
		if nodePtr != nil && *nodePtr != nil {
			val := (*nodePtr).(Alias)
			node.Alias = &val
		}
	}

	if fields["rtindex"] != nil {
		err = json.Unmarshal(fields["rtindex"], &node.Rtindex)
		if err != nil {
			return
		}
	}

	return
}
