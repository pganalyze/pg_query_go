// Auto-generated from postgres/src/include/nodes/parsenodes.h - DO NOT EDIT

package pg_query

import "encoding/json"

/* ----------------------
 *		Create {Aggregate|Operator|Type} Statement
 * ----------------------
 */
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
	var fields map[string]json.RawMessage

	err = json.Unmarshal(input, &fields)
	if err != nil {
		return
	}

	if fields["kind"] != nil {
		err = json.Unmarshal(fields["kind"], &node.Kind)
		if err != nil {
			return
		}
	}

	if fields["oldstyle"] != nil {
		err = json.Unmarshal(fields["oldstyle"], &node.Oldstyle)
		if err != nil {
			return
		}
	}

	if fields["defnames"] != nil {
		node.Defnames, err = UnmarshalNodeArrayJSON(fields["defnames"])
		if err != nil {
			return
		}
	}

	if fields["args"] != nil {
		node.Args, err = UnmarshalNodeArrayJSON(fields["args"])
		if err != nil {
			return
		}
	}

	if fields["definition"] != nil {
		node.Definition, err = UnmarshalNodeArrayJSON(fields["definition"])
		if err != nil {
			return
		}
	}

	return
}
