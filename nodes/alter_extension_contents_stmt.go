// Auto-generated from postgres/src/include/nodes/parsenodes.h - DO NOT EDIT

package pg_query

import "encoding/json"

/* ----------------------
 *		Create/Alter Extension Statements
 * ----------------------
 */
type AlterExtensionContentsStmt struct {
	Extname *string    `json:"extname"` /* Extension's name */
	Action  int        `json:"action"`  /* +1 = add object, -1 = drop object */
	Objtype ObjectType `json:"objtype"` /* Object's type */
	Objname List       `json:"objname"` /* Qualified name of the object */
	Objargs List       `json:"objargs"` /* Arguments if needed (eg, for functions) */
}

func (node AlterExtensionContentsStmt) MarshalJSON() ([]byte, error) {
	type AlterExtensionContentsStmtMarshalAlias AlterExtensionContentsStmt
	return json.Marshal(map[string]interface{}{
		"AlterExtensionContentsStmt": (*AlterExtensionContentsStmtMarshalAlias)(&node),
	})
}

func (node *AlterExtensionContentsStmt) UnmarshalJSON(input []byte) (err error) {
	var fields map[string]json.RawMessage

	err = json.Unmarshal(input, &fields)
	if err != nil {
		return
	}

	if fields["extname"] != nil {
		err = json.Unmarshal(fields["extname"], &node.Extname)
		if err != nil {
			return
		}
	}

	if fields["action"] != nil {
		err = json.Unmarshal(fields["action"], &node.Action)
		if err != nil {
			return
		}
	}

	if fields["objtype"] != nil {
		err = json.Unmarshal(fields["objtype"], &node.Objtype)
		if err != nil {
			return
		}
	}

	if fields["objname"] != nil {
		node.Objname.Items, err = UnmarshalNodeArrayJSON(fields["objname"])
		if err != nil {
			return
		}
	}

	if fields["objargs"] != nil {
		node.Objargs.Items, err = UnmarshalNodeArrayJSON(fields["objargs"])
		if err != nil {
			return
		}
	}

	return
}
