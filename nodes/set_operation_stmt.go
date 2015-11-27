// Auto-generated from postgres/src/include/nodes/parsenodes.h - DO NOT EDIT

package pg_query

import "encoding/json"

type SetOperationStmt struct {
	Op   SetOperation `json:"op"`   /* type of set op */
	All  bool         `json:"all"`  /* ALL specified? */
	Larg Node         `json:"larg"` /* left child */
	Rarg Node         `json:"rarg"` /* right child */

	/* Eventually add fields for CORRESPONDING spec here */

	/* Fields derived during parse analysis: */
	ColTypes      []Node `json:"colTypes"`      /* OID list of output column type OIDs */
	ColTypmods    []Node `json:"colTypmods"`    /* integer list of output column typmods */
	ColCollations []Node `json:"colCollations"` /* OID list of output column collation OIDs */
	GroupClauses  []Node `json:"groupClauses"`  /* a list of SortGroupClause's */

	/* groupClauses is NIL if UNION ALL, but must be set otherwise */
}

func (node SetOperationStmt) MarshalJSON() ([]byte, error) {
	type SetOperationStmtMarshalAlias SetOperationStmt
	return json.Marshal(map[string]interface{}{
		"SETOPERATIONSTMT": (*SetOperationStmtMarshalAlias)(&node),
	})
}

func (node *SetOperationStmt) UnmarshalJSON(input []byte) (err error) {
	var fields map[string]json.RawMessage

	err = json.Unmarshal(input, &fields)
	if err != nil {
		return
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

	if fields["colTypes"] != nil {
		node.ColTypes, err = UnmarshalNodeArrayJSON(fields["colTypes"])
		if err != nil {
			return
		}
	}

	if fields["colTypmods"] != nil {
		node.ColTypmods, err = UnmarshalNodeArrayJSON(fields["colTypmods"])
		if err != nil {
			return
		}
	}

	if fields["colCollations"] != nil {
		node.ColCollations, err = UnmarshalNodeArrayJSON(fields["colCollations"])
		if err != nil {
			return
		}
	}

	if fields["groupClauses"] != nil {
		node.GroupClauses, err = UnmarshalNodeArrayJSON(fields["groupClauses"])
		if err != nil {
			return
		}
	}

	return
}
