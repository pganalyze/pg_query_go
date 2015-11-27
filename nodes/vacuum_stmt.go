// Auto-generated from postgres/src/include/nodes/parsenodes.h - DO NOT EDIT

package pg_query

import "encoding/json"

/* ----------------------
 *		Vacuum and Analyze Statements
 *
 * Even though these are nominally two statements, it's convenient to use
 * just one node type for both.  Note that at least one of VACOPT_VACUUM
 * and VACOPT_ANALYZE must be set in options.  VACOPT_FREEZE is an internal
 * convenience for the grammar and is not examined at runtime --- the
 * freeze_min_age and freeze_table_age fields are what matter.
 * ----------------------
 */
type VacuumStmt struct {
	Options               int `json:"options"`                  /* OR of VacuumOption flags */
	FreezeMinAge          int `json:"freeze_min_age"`           /* min freeze age, or -1 to use default */
	FreezeTableAge        int `json:"freeze_table_age"`         /* age at which to scan whole table */
	MultixactFreezeMinAge int `json:"multixact_freeze_min_age"` /* min multixact freeze age,
	 * or -1 to use default */

	MultixactFreezeTableAge int `json:"multixact_freeze_table_age"` /* multixact age at which to
	 * scan whole table */

	Relation *RangeVar `json:"relation"` /* single table to process, or NULL */
	VaCols   []Node    `json:"va_cols"`  /* list of column names, or NIL for all */
}

func (node VacuumStmt) MarshalJSON() ([]byte, error) {
	type VacuumStmtMarshalAlias VacuumStmt
	return json.Marshal(map[string]interface{}{
		"VACUUM": (*VacuumStmtMarshalAlias)(&node),
	})
}

func (node *VacuumStmt) UnmarshalJSON(input []byte) (err error) {
	var fields map[string]json.RawMessage

	err = json.Unmarshal(input, &fields)
	if err != nil {
		return
	}

	if fields["options"] != nil {
		err = json.Unmarshal(fields["options"], &node.Options)
		if err != nil {
			return
		}
	}

	if fields["freeze_min_age"] != nil {
		err = json.Unmarshal(fields["freeze_min_age"], &node.FreezeMinAge)
		if err != nil {
			return
		}
	}

	if fields["freeze_table_age"] != nil {
		err = json.Unmarshal(fields["freeze_table_age"], &node.FreezeTableAge)
		if err != nil {
			return
		}
	}

	if fields["multixact_freeze_min_age"] != nil {
		err = json.Unmarshal(fields["multixact_freeze_min_age"], &node.MultixactFreezeMinAge)
		if err != nil {
			return
		}
	}

	if fields["multixact_freeze_table_age"] != nil {
		err = json.Unmarshal(fields["multixact_freeze_table_age"], &node.MultixactFreezeTableAge)
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

	if fields["va_cols"] != nil {
		node.VaCols, err = UnmarshalNodeArrayJSON(fields["va_cols"])
		if err != nil {
			return
		}
	}

	return
}
