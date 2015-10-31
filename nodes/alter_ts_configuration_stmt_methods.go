// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

func (node AlterTSConfigurationStmt) MarshalJSON() ([]byte, error) {
	type AlterTSConfigurationStmtMarshalAlias AlterTSConfigurationStmt
	return json.Marshal(map[string]interface{}{
		"ALTERTSCONFIGURATIONSTMT": (*AlterTSConfigurationStmtMarshalAlias)(&node),
	})
}

func (node *AlterTSConfigurationStmt) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}

func (node AlterTSConfigurationStmt) Deparse() string {
	panic("Not Implemented")
}
