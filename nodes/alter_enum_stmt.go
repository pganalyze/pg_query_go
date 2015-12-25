// Auto-generated from postgres/src/include/nodes/parsenodes.h - DO NOT EDIT

package pg_query

import "encoding/json"

/* ----------------------
 *		Alter Type Statement, enum types
 * ----------------------
 */
type AlterEnumStmt struct {
	TypeName       List    `json:"typeName"`       /* qualified name (list of Value strings) */
	NewVal         *string `json:"newVal"`         /* new enum value's name */
	NewValNeighbor *string `json:"newValNeighbor"` /* neighboring enum value, if specified */
	NewValIsAfter  bool    `json:"newValIsAfter"`  /* place new enum value after neighbor? */
	SkipIfExists   bool    `json:"skipIfExists"`   /* no error if label already exists */
}

func (node AlterEnumStmt) MarshalJSON() ([]byte, error) {
	type AlterEnumStmtMarshalAlias AlterEnumStmt
	return json.Marshal(map[string]interface{}{
		"AlterEnumStmt": (*AlterEnumStmtMarshalAlias)(&node),
	})
}

func (node *AlterEnumStmt) UnmarshalJSON(input []byte) (err error) {
	var fields map[string]json.RawMessage

	err = json.Unmarshal(input, &fields)
	if err != nil {
		return
	}

	if fields["typeName"] != nil {
		node.TypeName.Items, err = UnmarshalNodeArrayJSON(fields["typeName"])
		if err != nil {
			return
		}
	}

	if fields["newVal"] != nil {
		err = json.Unmarshal(fields["newVal"], &node.NewVal)
		if err != nil {
			return
		}
	}

	if fields["newValNeighbor"] != nil {
		err = json.Unmarshal(fields["newValNeighbor"], &node.NewValNeighbor)
		if err != nil {
			return
		}
	}

	if fields["newValIsAfter"] != nil {
		err = json.Unmarshal(fields["newValIsAfter"], &node.NewValIsAfter)
		if err != nil {
			return
		}
	}

	if fields["skipIfExists"] != nil {
		err = json.Unmarshal(fields["skipIfExists"], &node.SkipIfExists)
		if err != nil {
			return
		}
	}

	return
}
