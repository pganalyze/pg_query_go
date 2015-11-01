// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

type AlterObjectSchemaStmt struct {
	ObjectType ObjectType `json:"objectType"` /* OBJECT_TABLE, OBJECT_TYPE, etc */
	Relation   *RangeVar  `json:"relation"`   /* in case it's a table */
	Object     []Node     `json:"object"`     /* in case it's some other object */
	Objarg     []Node     `json:"objarg"`     /* argument types, if applicable */
	Newschema  *string    `json:"newschema"`  /* the new schema */
	MissingOk  bool       `json:"missing_ok"` /* skip error if missing? */
}

func (node AlterObjectSchemaStmt) MarshalJSON() ([]byte, error) {
	type AlterObjectSchemaStmtMarshalAlias AlterObjectSchemaStmt
	return json.Marshal(map[string]interface{}{
		"ALTEROBJECTSCHEMASTMT": (*AlterObjectSchemaStmtMarshalAlias)(&node),
	})
}

func (node *AlterObjectSchemaStmt) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}
