// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

type AlterFunctionStmt struct {
	Func    *FuncWithArgs `json:"func"`    /* name and args of function */
	Actions []Node        `json:"actions"` /* list of DefElem */
}

func (node AlterFunctionStmt) MarshalJSON() ([]byte, error) {
	type AlterFunctionStmtMarshalAlias AlterFunctionStmt
	return json.Marshal(map[string]interface{}{
		"ALTERFUNCTIONSTMT": (*AlterFunctionStmtMarshalAlias)(&node),
	})
}

func (node *AlterFunctionStmt) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}
