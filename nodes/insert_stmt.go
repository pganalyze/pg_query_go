// Auto-generated - DO NOT EDIT

package pg_query

type InsertStmt struct {
	Relation      *RangeVar   `json:"relation"`      /* relation to insert into */
	Cols          []Node      `json:"cols"`          /* optional: names of the target columns */
	SelectStmt    Node        `json:"selectStmt"`    /* the source SELECT/VALUES, or NULL */
	ReturningList []Node      `json:"returningList"` /* list of expressions to return */
	WithClause    *WithClause `json:"withClause"`    /* WITH clause */
}
