// Auto-generated - DO NOT EDIT

package pg_query

type AlterFunctionStmt struct {
	Func    *FuncWithArgs `json:"func"`    /* name and args of function */
	Actions []Node        `json:"actions"` /* list of DefElem */
}
