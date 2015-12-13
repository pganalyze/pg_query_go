// Auto-generated from postgres/src/include/nodes/parsenodes.h - DO NOT EDIT

package pg_query

import "encoding/json"

/* ----------------------
 *		Create Type Statement, range types
 * ----------------------
 */
type CreateRangeStmt struct {
	TypeName []Node `json:"typeName"` /* qualified name (list of Value strings) */
	Params   []Node `json:"params"`   /* range parameters (list of DefElem) */
}

func (node CreateRangeStmt) MarshalJSON() ([]byte, error) {
	type CreateRangeStmtMarshalAlias CreateRangeStmt
	return json.Marshal(map[string]interface{}{
		"CreateRangeStmt": (*CreateRangeStmtMarshalAlias)(&node),
	})
}

func (node *CreateRangeStmt) UnmarshalJSON(input []byte) (err error) {
	var fields map[string]json.RawMessage

	err = json.Unmarshal(input, &fields)
	if err != nil {
		return
	}

	if fields["typeName"] != nil {
		node.TypeName, err = UnmarshalNodeArrayJSON(fields["typeName"])
		if err != nil {
			return
		}
	}

	if fields["params"] != nil {
		node.Params, err = UnmarshalNodeArrayJSON(fields["params"])
		if err != nil {
			return
		}
	}

	return
}
