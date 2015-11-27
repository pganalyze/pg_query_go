// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

type CreateTableAsStmt struct {
	Query        Node        `json:"query"`          /* the query (see comments above) */
	Into         *IntoClause `json:"into"`           /* destination table */
	Relkind      ObjectType  `json:"relkind"`        /* OBJECT_TABLE or OBJECT_MATVIEW */
	IsSelectInto bool        `json:"is_select_into"` /* it was written as SELECT INTO */
}

func (node CreateTableAsStmt) MarshalJSON() ([]byte, error) {
	type CreateTableAsStmtMarshalAlias CreateTableAsStmt
	return json.Marshal(map[string]interface{}{
		"CREATE TABLE AS": (*CreateTableAsStmtMarshalAlias)(&node),
	})
}

func (node *CreateTableAsStmt) UnmarshalJSON(input []byte) (err error) {
	var fields map[string]json.RawMessage

	err = json.Unmarshal(input, &fields)
	if err != nil {
		return
	}

	if fields["query"] != nil {
		node.Query, err = UnmarshalNodeJSON(fields["query"])
		if err != nil {
			return
		}
	}

	if fields["into"] != nil {
		var nodePtr *Node
		nodePtr, err = UnmarshalNodePtrJSON(fields["into"])
		if err != nil {
			return
		}
		if nodePtr != nil && *nodePtr != nil {
			val := (*nodePtr).(IntoClause)
			node.Into = &val
		}
	}

	if fields["relkind"] != nil {
		err = json.Unmarshal(fields["relkind"], &node.Relkind)
		if err != nil {
			return
		}
	}

	if fields["is_select_into"] != nil {
		err = json.Unmarshal(fields["is_select_into"], &node.IsSelectInto)
		if err != nil {
			return
		}
	}

	return
}
