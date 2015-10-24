// Auto-generated - DO NOT EDIT

package pg_query

type DeleteStmt struct {
	Relation      *RangeVar   `json:"relation"`      /* relation to delete from */
	UsingClause   []Node      `json:"usingClause"`   /* optional using clause for more tables */
	WhereClause   Node        `json:"whereClause"`   /* qualifications */
	ReturningList []Node      `json:"returningList"` /* list of expressions to return */
	WithClause    *WithClause `json:"withClause"`    /* WITH clause */
}
