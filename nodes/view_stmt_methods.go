// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

func (node ViewStmt) MarshalJSON() ([]byte, error) {
	type ViewStmtMarshalAlias ViewStmt
	return json.Marshal(map[string]interface{}{
		"VIEWSTMT": (*ViewStmtMarshalAlias)(&node),
	})
}

func (node *ViewStmt) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}

func (node ViewStmt) Deparse() string {
	panic("Not Implemented")
}
