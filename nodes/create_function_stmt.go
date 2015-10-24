// Auto-generated - DO NOT EDIT

package pg_query

type CreateFunctionStmt struct {
	Replace    bool      `json:"replace"`    /* T => replace if already exists */
	Funcname   []Node    `json:"funcname"`   /* qualified name of function to create */
	Parameters []Node    `json:"parameters"` /* a list of FunctionParameter */
	ReturnType *TypeName `json:"returnType"` /* the return type */
	Options    []Node    `json:"options"`    /* a list of DefElem */
	WithClause []Node    `json:"withClause"` /* a list of DefElem */
}
