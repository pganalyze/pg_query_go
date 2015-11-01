// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

type CompositeTypeStmt struct {
	Typevar    *RangeVar `json:"typevar"`    /* the composite type to be created */
	Coldeflist []Node    `json:"coldeflist"` /* list of ColumnDef nodes */
}

func (node CompositeTypeStmt) MarshalJSON() ([]byte, error) {
	type CompositeTypeStmtMarshalAlias CompositeTypeStmt
	return json.Marshal(map[string]interface{}{
		"COMPOSITETYPESTMT": (*CompositeTypeStmtMarshalAlias)(&node),
	})
}

func (node *CompositeTypeStmt) UnmarshalJSON(input []byte) (err error) {
	err = UnmarshalNodeFieldJSON(input, node)
	return
}
