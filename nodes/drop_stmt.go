// Auto-generated - DO NOT EDIT

package pg_query

type DropStmt struct {
	Objects    []Node       `json:"objects"`    /* list of sublists of names (as Values) */
	Arguments  []Node       `json:"arguments"`  /* list of sublists of arguments (as Values) */
	RemoveType ObjectType   `json:"removeType"` /* object type */
	Behavior   DropBehavior `json:"behavior"`   /* RESTRICT or CASCADE behavior */
	MissingOk  bool         `json:"missing_ok"` /* skip error if object is missing? */
	Concurrent bool         `json:"concurrent"` /* drop index concurrently? */
}
