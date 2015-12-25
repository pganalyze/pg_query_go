// Auto-generated from postgres/src/include/nodes/parsenodes.h - DO NOT EDIT

package pg_query

import "encoding/json"

/* ----------------------
 *				SECURITY LABEL Statement
 * ----------------------
 */
type SecLabelStmt struct {
	Objtype  ObjectType `json:"objtype"`  /* Object's type */
	Objname  List       `json:"objname"`  /* Qualified name of the object */
	Objargs  List       `json:"objargs"`  /* Arguments if needed (eg, for functions) */
	Provider *string    `json:"provider"` /* Label provider (or NULL) */
	Label    *string    `json:"label"`    /* New security label to be assigned */
}

func (node SecLabelStmt) MarshalJSON() ([]byte, error) {
	type SecLabelStmtMarshalAlias SecLabelStmt
	return json.Marshal(map[string]interface{}{
		"SecLabelStmt": (*SecLabelStmtMarshalAlias)(&node),
	})
}

func (node *SecLabelStmt) UnmarshalJSON(input []byte) (err error) {
	var fields map[string]json.RawMessage

	err = json.Unmarshal(input, &fields)
	if err != nil {
		return
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

	if fields["provider"] != nil {
		err = json.Unmarshal(fields["provider"], &node.Provider)
		if err != nil {
			return
		}
	}

	if fields["label"] != nil {
		err = json.Unmarshal(fields["label"], &node.Label)
		if err != nil {
			return
		}
	}

	return
}
