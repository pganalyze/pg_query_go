// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

type UpdateStmt struct {
	Relation      *RangeVar   `json:"relation"`      /* relation to update */
	TargetList    []Node      `json:"targetList"`    /* the target list (of ResTarget) */
	WhereClause   Node        `json:"whereClause"`   /* qualifications */
	FromClause    []Node      `json:"fromClause"`    /* optional from clause for more tables */
	ReturningList []Node      `json:"returningList"` /* list of expressions to return */
	WithClause    *WithClause `json:"withClause"`    /* WITH clause */
}

func (node UpdateStmt) MarshalJSON() ([]byte, error) {
	type UpdateStmtMarshalAlias UpdateStmt
	return json.Marshal(map[string]interface{}{
		"UPDATE": (*UpdateStmtMarshalAlias)(&node),
	})
}

func (node *UpdateStmt) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}
