// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

type AlterExtensionContentsStmt struct {
	Extname *string    `json:"extname"` /* Extension's name */
	Action  int        `json:"action"`  /* +1 = add object, -1 = drop object */
	Objtype ObjectType `json:"objtype"` /* Object's type */
	Objname []Node     `json:"objname"` /* Qualified name of the object */
	Objargs []Node     `json:"objargs"` /* Arguments if needed (eg, for functions) */
}

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
