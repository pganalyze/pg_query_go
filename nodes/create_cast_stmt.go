// Auto-generated - DO NOT EDIT

package pg_query

type CreateCastStmt struct {
	Sourcetype *TypeName       `json:"sourcetype"`
	Targettype *TypeName       `json:"targettype"`
	Func       *FuncWithArgs   `json:"func"`
	Context    CoercionContext `json:"context"`
	Inout      bool            `json:"inout"`
}
