// Auto-generated - DO NOT EDIT

package pg_query

type PrepareStmt struct {
	Name     *string `json:"name"`     /* Name of plan, arbitrary */
	Argtypes []Node  `json:"argtypes"` /* Types of parameters (List of TypeName) */
	Query    Node    `json:"query"`    /* The query itself (as a raw parsetree) */
}
