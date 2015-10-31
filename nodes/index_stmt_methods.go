// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

func (node IndexStmt) MarshalJSON() ([]byte, error) {
	type IndexStmtMarshalAlias IndexStmt
	return json.Marshal(map[string]interface{}{
		"INDEXSTMT": (*IndexStmtMarshalAlias)(&node),
	})
}

func (node *IndexStmt) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}

func (node IndexStmt) Deparse() string {
	panic("Not Implemented")
}
