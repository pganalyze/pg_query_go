package pg_query

import "encoding/json"

func (selectStmt SelectStmt) MarshalJSON() ([]byte, error) {
	type SelectStmtMarshalAlias SelectStmt
	return json.Marshal(map[string]interface{}{
		"SELECT": (*SelectStmtMarshalAlias)(&selectStmt),
	})
}

func (selectStmt *SelectStmt) UnmarshalJSON(input []byte) error {
	return UnmarshalNodeFieldJSON(input, selectStmt)
}

func (selectStmt SelectStmt) Deparse() string {
	panic("Not Implemented")
}
