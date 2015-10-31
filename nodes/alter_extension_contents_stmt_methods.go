// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

func (node AlterExtensionContentsStmt) MarshalJSON() ([]byte, error) {
	type AlterExtensionContentsStmtMarshalAlias AlterExtensionContentsStmt
	return json.Marshal(map[string]interface{}{
		"ALTEREXTENSIONCONTENTSSTMT": (*AlterExtensionContentsStmtMarshalAlias)(&node),
	})
}

func (node *AlterExtensionContentsStmt) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}

func (node AlterExtensionContentsStmt) Deparse() string {
	panic("Not Implemented")
}
