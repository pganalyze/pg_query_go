// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

type SecLabelStmt struct {
	Objtype  ObjectType `json:"objtype"`  /* Object's type */
	Objname  []Node     `json:"objname"`  /* Qualified name of the object */
	Objargs  []Node     `json:"objargs"`  /* Arguments if needed (eg, for functions) */
	Provider *string    `json:"provider"` /* Label provider (or NULL) */
	Label    *string    `json:"label"`    /* New security label to be assigned */
}

func (node SecLabelStmt) MarshalJSON() ([]byte, error) {
	type SecLabelStmtMarshalAlias SecLabelStmt
	return json.Marshal(map[string]interface{}{
		"SECLABELSTMT": (*SecLabelStmtMarshalAlias)(&node),
	})
}

func (node *SecLabelStmt) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}
