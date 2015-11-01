// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

type AlterOwnerStmt struct {
	ObjectType ObjectType `json:"objectType"` /* OBJECT_TABLE, OBJECT_TYPE, etc */
	Relation   *RangeVar  `json:"relation"`   /* in case it's a table */
	Object     []Node     `json:"object"`     /* in case it's some other object */
	Objarg     []Node     `json:"objarg"`     /* argument types, if applicable */
	Newowner   *string    `json:"newowner"`   /* the new owner */
}

func (node AlterOwnerStmt) MarshalJSON() ([]byte, error) {
	type AlterOwnerStmtMarshalAlias AlterOwnerStmt
	return json.Marshal(map[string]interface{}{
		"ALTEROWNERSTMT": (*AlterOwnerStmtMarshalAlias)(&node),
	})
}

func (node *AlterOwnerStmt) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}
