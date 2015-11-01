// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

type DropStmt struct {
	Objects    []Node       `json:"objects"`    /* list of sublists of names (as Values) */
	Arguments  []Node       `json:"arguments"`  /* list of sublists of arguments (as Values) */
	RemoveType ObjectType   `json:"removeType"` /* object type */
	Behavior   DropBehavior `json:"behavior"`   /* RESTRICT or CASCADE behavior */
	MissingOk  bool         `json:"missing_ok"` /* skip error if object is missing? */
	Concurrent bool         `json:"concurrent"` /* drop index concurrently? */
}

func (node DropStmt) MarshalJSON() ([]byte, error) {
	type DropStmtMarshalAlias DropStmt
	return json.Marshal(map[string]interface{}{
		"DROP": (*DropStmtMarshalAlias)(&node),
	})
}

func (node *DropStmt) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}
