// Auto-generated from postgres/src/include/nodes/parsenodes.h - DO NOT EDIT

package pg_query

import "encoding/json"

type CreateFunctionStmt struct {
	Replace    bool      `json:"replace"`    /* T => replace if already exists */
	Funcname   []Node    `json:"funcname"`   /* qualified name of function to create */
	Parameters []Node    `json:"parameters"` /* a list of FunctionParameter */
	ReturnType *TypeName `json:"returnType"` /* the return type */
	Options    []Node    `json:"options"`    /* a list of DefElem */
	WithClause []Node    `json:"withClause"` /* a list of DefElem */
}

func (node CreateFunctionStmt) MarshalJSON() ([]byte, error) {
	type CreateFunctionStmtMarshalAlias CreateFunctionStmt
	return json.Marshal(map[string]interface{}{
		"CREATEFUNCTIONSTMT": (*CreateFunctionStmtMarshalAlias)(&node),
	})
}

func (node *CreateFunctionStmt) UnmarshalJSON(input []byte) (err error) {
	var fields map[string]json.RawMessage

	err = json.Unmarshal(input, &fields)
	if err != nil {
		return
	}

	if fields["replace"] != nil {
		err = json.Unmarshal(fields["replace"], &node.Replace)
		if err != nil {
			return
		}
	}

	if fields["funcname"] != nil {
		node.Funcname, err = UnmarshalNodeArrayJSON(fields["funcname"])
		if err != nil {
			return
		}
	}

	if fields["parameters"] != nil {
		node.Parameters, err = UnmarshalNodeArrayJSON(fields["parameters"])
		if err != nil {
			return
		}
	}

	if fields["returnType"] != nil {
		var nodePtr *Node
		nodePtr, err = UnmarshalNodePtrJSON(fields["returnType"])
		if err != nil {
			return
		}
		if nodePtr != nil && *nodePtr != nil {
			val := (*nodePtr).(TypeName)
			node.ReturnType = &val
		}
	}

	if fields["options"] != nil {
		node.Options, err = UnmarshalNodeArrayJSON(fields["options"])
		if err != nil {
			return
		}
	}

	if fields["withClause"] != nil {
		node.WithClause, err = UnmarshalNodeArrayJSON(fields["withClause"])
		if err != nil {
			return
		}
	}

	return
}
