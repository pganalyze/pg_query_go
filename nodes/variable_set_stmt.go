// Auto-generated - DO NOT EDIT

package pg_query

type VariableSetStmt struct {
	Kind    VariableSetKind `json:"kind"`
	Name    *string         `json:"name"`     /* variable to be set */
	Args    []Node          `json:"args"`     /* List of A_Const nodes */
	IsLocal bool            `json:"is_local"` /* SET LOCAL? */
}
