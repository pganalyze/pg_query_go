// Auto-generated - DO NOT EDIT

package pg_query

type AlterEnumStmt struct {
	TypeName       []Node  `json:"typeName"`       /* qualified name (list of Value strings) */
	NewVal         *string `json:"newVal"`         /* new enum value's name */
	NewValNeighbor *string `json:"newValNeighbor"` /* neighboring enum value, if specified */
	NewValIsAfter  bool    `json:"newValIsAfter"`  /* place new enum value after neighbor? */
	SkipIfExists   bool    `json:"skipIfExists"`   /* no error if label already exists */
}
