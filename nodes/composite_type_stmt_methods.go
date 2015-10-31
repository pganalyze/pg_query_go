// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

func (node CompositeTypeStmt) MarshalJSON() ([]byte, error) {
	type CompositeTypeStmtMarshalAlias CompositeTypeStmt
	return json.Marshal(map[string]interface{}{
		"COMPOSITETYPESTMT": (*CompositeTypeStmtMarshalAlias)(&node),
	})
}

func (node *CompositeTypeStmt) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}

func (node CompositeTypeStmt) Deparse() string {
	panic("Not Implemented")
}
