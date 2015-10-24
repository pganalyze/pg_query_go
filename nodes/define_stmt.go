// Auto-generated - DO NOT EDIT

package pg_query

type DefineStmt struct {
	Kind       ObjectType `json:"kind"`       /* aggregate, operator, type */
	Oldstyle   bool       `json:"oldstyle"`   /* hack to signal old CREATE AGG syntax */
	Defnames   []Node     `json:"defnames"`   /* qualified name (list of Value strings) */
	Args       []Node     `json:"args"`       /* a list of TypeName (if needed) */
	Definition []Node     `json:"definition"` /* a list of DefElem */
}
