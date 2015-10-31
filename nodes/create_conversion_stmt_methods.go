// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

func (node CreateConversionStmt) MarshalJSON() ([]byte, error) {
	type CreateConversionStmtMarshalAlias CreateConversionStmt
	return json.Marshal(map[string]interface{}{
		"CREATECONVERSIONSTMT": (*CreateConversionStmtMarshalAlias)(&node),
	})
}

func (node *CreateConversionStmt) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}

func (node CreateConversionStmt) Deparse() string {
	panic("Not Implemented")
}
