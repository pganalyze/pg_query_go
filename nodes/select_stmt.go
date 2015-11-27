// Auto-generated from postgres/src/include/nodes/parsenodes.h - DO NOT EDIT

package pg_query

import "encoding/json"

type SelectStmt struct {
	/*
	 * These fields are used only in "leaf" SelectStmts.
	 */
	DistinctClause []Node `json:"distinctClause"` /* NULL, list of DISTINCT ON exprs, or
	 * lcons(NIL,NIL) for all (SELECT DISTINCT) */

	IntoClause   *IntoClause `json:"intoClause"`   /* target for SELECT INTO */
	TargetList   []Node      `json:"targetList"`   /* the target list (of ResTarget) */
	FromClause   []Node      `json:"fromClause"`   /* the FROM clause */
	WhereClause  Node        `json:"whereClause"`  /* WHERE qualification */
	GroupClause  []Node      `json:"groupClause"`  /* GROUP BY clauses */
	HavingClause Node        `json:"havingClause"` /* HAVING conditional-expression */
	WindowClause []Node      `json:"windowClause"` /* WINDOW window_name AS (...), ... */

	/*
	 * In a "leaf" node representing a VALUES list, the above fields are all
	 * null, and instead this field is set.  Note that the elements of the
	 * sublists are just expressions, without ResTarget decoration. Also note
	 * that a list element can be DEFAULT (represented as a SetToDefault
	 * node), regardless of the context of the VALUES list. It's up to parse
	 * analysis to reject that where not valid.
	 */
	ValuesLists [][]Node `json:"valuesLists"` /* untransformed list of expression lists */

	/*
	 * These fields are used in both "leaf" SelectStmts and upper-level
	 * SelectStmts.
	 */
	SortClause    []Node      `json:"sortClause"`    /* sort clause (a list of SortBy's) */
	LimitOffset   Node        `json:"limitOffset"`   /* # of result tuples to skip */
	LimitCount    Node        `json:"limitCount"`    /* # of result tuples to return */
	LockingClause []Node      `json:"lockingClause"` /* FOR UPDATE (list of LockingClause's) */
	WithClause    *WithClause `json:"withClause"`    /* WITH clause */

	/*
	 * These fields are used only in upper-level SelectStmts.
	 */
	Op   SetOperation `json:"op"`   /* type of set op */
	All  bool         `json:"all"`  /* ALL specified? */
	Larg *SelectStmt  `json:"larg"` /* left child */
	Rarg *SelectStmt  `json:"rarg"` /* right child */

	/* Eventually add fields for CORRESPONDING spec here */
}

func (node SelectStmt) MarshalJSON() ([]byte, error) {
	type SelectStmtMarshalAlias SelectStmt
	return json.Marshal(map[string]interface{}{
		"SELECT": (*SelectStmtMarshalAlias)(&node),
	})
}

func (node *SelectStmt) UnmarshalJSON(input []byte) (err error) {
	var fields map[string]json.RawMessage

	err = json.Unmarshal(input, &fields)
	if err != nil {
		return
	}

	if fields["distinctClause"] != nil {
		node.DistinctClause, err = UnmarshalNodeArrayJSON(fields["distinctClause"])
		if err != nil {
			return
		}
	}

	if fields["intoClause"] != nil {
		var nodePtr *Node
		nodePtr, err = UnmarshalNodePtrJSON(fields["intoClause"])
		if err != nil {
			return
		}
		if nodePtr != nil && *nodePtr != nil {
			val := (*nodePtr).(IntoClause)
			node.IntoClause = &val
		}
	}

	if fields["targetList"] != nil {
		node.TargetList, err = UnmarshalNodeArrayJSON(fields["targetList"])
		if err != nil {
			return
		}
	}

	if fields["fromClause"] != nil {
		node.FromClause, err = UnmarshalNodeArrayJSON(fields["fromClause"])
		if err != nil {
			return
		}
	}

	if fields["whereClause"] != nil {
		node.WhereClause, err = UnmarshalNodeJSON(fields["whereClause"])
		if err != nil {
			return
		}
	}

	if fields["groupClause"] != nil {
		node.GroupClause, err = UnmarshalNodeArrayJSON(fields["groupClause"])
		if err != nil {
			return
		}
	}

	if fields["havingClause"] != nil {
		node.HavingClause, err = UnmarshalNodeJSON(fields["havingClause"])
		if err != nil {
			return
		}
	}

	if fields["windowClause"] != nil {
		node.WindowClause, err = UnmarshalNodeArrayJSON(fields["windowClause"])
		if err != nil {
			return
		}
	}

	if fields["valuesLists"] != nil {
		node.ValuesLists, err = UnmarshalNodeArrayArrayJSON(fields["valuesLists"])
		if err != nil {
			return
		}
	}

	if fields["sortClause"] != nil {
		node.SortClause, err = UnmarshalNodeArrayJSON(fields["sortClause"])
		if err != nil {
			return
		}
	}

	if fields["limitOffset"] != nil {
		node.LimitOffset, err = UnmarshalNodeJSON(fields["limitOffset"])
		if err != nil {
			return
		}
	}

	if fields["limitCount"] != nil {
		node.LimitCount, err = UnmarshalNodeJSON(fields["limitCount"])
		if err != nil {
			return
		}
	}

	if fields["lockingClause"] != nil {
		node.LockingClause, err = UnmarshalNodeArrayJSON(fields["lockingClause"])
		if err != nil {
			return
		}
	}

	if fields["withClause"] != nil {
		var nodePtr *Node
		nodePtr, err = UnmarshalNodePtrJSON(fields["withClause"])
		if err != nil {
			return
		}
		if nodePtr != nil && *nodePtr != nil {
			val := (*nodePtr).(WithClause)
			node.WithClause = &val
		}
	}

	if fields["op"] != nil {
		err = json.Unmarshal(fields["op"], &node.Op)
		if err != nil {
			return
		}
	}

	if fields["all"] != nil {
		err = json.Unmarshal(fields["all"], &node.All)
		if err != nil {
			return
		}
	}

	if fields["larg"] != nil {
		var nodePtr *Node
		nodePtr, err = UnmarshalNodePtrJSON(fields["larg"])
		if err != nil {
			return
		}
		if nodePtr != nil && *nodePtr != nil {
			val := (*nodePtr).(SelectStmt)
			node.Larg = &val
		}
	}

	if fields["rarg"] != nil {
		var nodePtr *Node
		nodePtr, err = UnmarshalNodePtrJSON(fields["rarg"])
		if err != nil {
			return
		}
		if nodePtr != nil && *nodePtr != nil {
			val := (*nodePtr).(SelectStmt)
			node.Rarg = &val
		}
	}

	return
}
