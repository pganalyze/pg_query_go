// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

type DefineStmt struct {
	Kind       ObjectType `json:"kind"`       /* aggregate, operator, type */
	Oldstyle   bool       `json:"oldstyle"`   /* hack to signal old CREATE AGG syntax */
	Defnames   []Node     `json:"defnames"`   /* qualified name (list of Value strings) */
	Args       []Node     `json:"args"`       /* a list of TypeName (if needed) */
	Definition []Node     `json:"definition"` /* a list of DefElem */
}

func (node DefineStmt) MarshalJSON() ([]byte, error) {
	type DefineStmtMarshalAlias DefineStmt
	return json.Marshal(map[string]interface{}{
		"DEFINESTMT": (*DefineStmtMarshalAlias)(&node),
	})
}

func (node *DefineStmt) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}
