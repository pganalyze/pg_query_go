package pg_query

type DeparseContext string

const (
	DeparseContextNone     DeparseContext = "None"
	DeparseContextAConst   DeparseContext = "A_Const"
	DeparseContextOperator DeparseContext = "Operator"
	DeparseContextSelect   DeparseContext = "Select"
	DeparseContextUpdate   DeparseContext = "Update"
	DeparseContextTypeName DeparseContext = "TypeName"
	DeparseContextFuncCall DeparseContext = "FuncCall"
)
