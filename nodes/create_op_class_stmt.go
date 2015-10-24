// Auto-generated - DO NOT EDIT

package pg_query

type CreateOpClassStmt struct {
	Opclassname  []Node    `json:"opclassname"`  /* qualified name (list of Value strings) */
	Opfamilyname []Node    `json:"opfamilyname"` /* qualified name (ditto); NIL if omitted */
	Amname       *string   `json:"amname"`       /* name of index AM opclass is for */
	Datatype     *TypeName `json:"datatype"`     /* datatype of indexed column */
	Items        []Node    `json:"items"`        /* List of CreateOpClassItem nodes */
	IsDefault    bool      `json:"isDefault"`    /* Should be marked as default for type? */
}
