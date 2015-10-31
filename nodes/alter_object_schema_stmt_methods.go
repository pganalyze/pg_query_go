// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

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

func (node AlterObjectSchemaStmt) Deparse() string {
	panic("Not Implemented")
}
