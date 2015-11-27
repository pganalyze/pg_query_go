// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

type ReindexStmt struct {
	Kind     ObjectType `json:"kind"`      /* OBJECT_INDEX, OBJECT_TABLE, etc. */
	Relation *RangeVar  `json:"relation"`  /* Table or index to reindex */
	Name     *string    `json:"name"`      /* name of database to reindex */
	DoSystem bool       `json:"do_system"` /* include system tables in database case */
	DoUser   bool       `json:"do_user"`   /* include user tables in database case */
}

func (node ReindexStmt) MarshalJSON() ([]byte, error) {
	type ReindexStmtMarshalAlias ReindexStmt
	return json.Marshal(map[string]interface{}{
		"REINDEXSTMT": (*ReindexStmtMarshalAlias)(&node),
	})
}

func (node *ReindexStmt) UnmarshalJSON(input []byte) (err error) {
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

	if fields["relation"] != nil {
		var nodePtr *Node
		nodePtr, err = UnmarshalNodePtrJSON(fields["relation"])
		if err != nil {
			return
		}
		if nodePtr != nil && *nodePtr != nil {
			val := (*nodePtr).(RangeVar)
			node.Relation = &val
		}
	}

	if fields["name"] != nil {
		err = json.Unmarshal(fields["name"], &node.Name)
		if err != nil {
			return
		}
	}

	if fields["do_system"] != nil {
		err = json.Unmarshal(fields["do_system"], &node.DoSystem)
		if err != nil {
			return
		}
	}

	if fields["do_user"] != nil {
		err = json.Unmarshal(fields["do_user"], &node.DoUser)
		if err != nil {
			return
		}
	}

	return
}
