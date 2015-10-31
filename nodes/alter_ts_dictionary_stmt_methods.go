// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

func (node AlterTSDictionaryStmt) MarshalJSON() ([]byte, error) {
	type AlterTSDictionaryStmtMarshalAlias AlterTSDictionaryStmt
	return json.Marshal(map[string]interface{}{
		"ALTERTSDICTIONARYSTMT": (*AlterTSDictionaryStmtMarshalAlias)(&node),
	})
}

func (node *AlterTSDictionaryStmt) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}

func (node AlterTSDictionaryStmt) Deparse() string {
	panic("Not Implemented")
}
