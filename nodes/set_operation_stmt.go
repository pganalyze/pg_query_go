// Auto-generated - DO NOT EDIT

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
	err = UnmarshalNodeFieldJSON(input, node)
	return
}
