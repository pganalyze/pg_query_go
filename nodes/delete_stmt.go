// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

type DeleteStmt struct {
	Relation      *RangeVar   `json:"relation"`      /* relation to delete from */
	UsingClause   []Node      `json:"usingClause"`   /* optional using clause for more tables */
	WhereClause   Node        `json:"whereClause"`   /* qualifications */
	ReturningList []Node      `json:"returningList"` /* list of expressions to return */
	WithClause    *WithClause `json:"withClause"`    /* WITH clause */
}

func (node DeleteStmt) MarshalJSON() ([]byte, error) {
	type DeleteStmtMarshalAlias DeleteStmt
	return json.Marshal(map[string]interface{}{
		"DELETE FROM": (*DeleteStmtMarshalAlias)(&node),
	})
}

func (node *DeleteStmt) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}
