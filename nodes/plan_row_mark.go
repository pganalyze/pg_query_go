// Auto-generated from postgres/src/include/nodes/plannodes.h - DO NOT EDIT

package pg_query

import "encoding/json"

/*
 * PlanRowMark -
 *	   plan-time representation of FOR [KEY] UPDATE/SHARE clauses
 *
 * When doing UPDATE, DELETE, or SELECT FOR UPDATE/SHARE, we create a separate
 * PlanRowMark node for each non-target relation in the query.  Relations that
 * are not specified as FOR UPDATE/SHARE are marked ROW_MARK_REFERENCE (if
 * regular tables) or ROW_MARK_COPY (if not).
 *
 * Initially all PlanRowMarks have rti == prti and isParent == false.
 * When the planner discovers that a relation is the root of an inheritance
 * tree, it sets isParent true, and adds an additional PlanRowMark to the
 * list for each child relation (including the target rel itself in its role
 * as a child).  The child entries have rti == child rel's RT index and
 * prti == parent's RT index, and can therefore be recognized as children by
 * the fact that prti != rti.
 *
 * The planner also adds resjunk output columns to the plan that carry
 * information sufficient to identify the locked or fetched rows.  For
 * regular tables (markType != ROW_MARK_COPY), these columns are named
 *		tableoid%u			OID of table
 *		ctid%u				TID of row
 * The tableoid column is only present for an inheritance hierarchy.
 * When markType == ROW_MARK_COPY, there is instead a single column named
 *		wholerow%u			whole-row value of relation
 * In all three cases, %u represents the rowmark ID number (rowmarkId).
 * This number is unique within a plan tree, except that child relation
 * entries copy their parent's rowmarkId.  (Assigning unique numbers
 * means we needn't renumber rowmarkIds when flattening subqueries, which
 * would require finding and renaming the resjunk columns as well.)
 * Note this means that all tables in an inheritance hierarchy share the
 * same resjunk column names.  However, in an inherited UPDATE/DELETE the
 * columns could have different physical column numbers in each subplan.
 */
type PlanRowMark struct {
	Rti       Index       `json:"rti"`       /* range table index of markable relation */
	Prti      Index       `json:"prti"`      /* range table index of parent relation */
	RowmarkId Index       `json:"rowmarkId"` /* unique identifier for resjunk columns */
	MarkType  RowMarkType `json:"markType"`  /* see enum above */
	NoWait    bool        `json:"noWait"`    /* NOWAIT option */
	IsParent  bool        `json:"isParent"`  /* true if this is a "dummy" parent entry */
}

func (node PlanRowMark) MarshalJSON() ([]byte, error) {
	type PlanRowMarkMarshalAlias PlanRowMark
	return json.Marshal(map[string]interface{}{
		"PLANROWMARK": (*PlanRowMarkMarshalAlias)(&node),
	})
}

func (node *PlanRowMark) UnmarshalJSON(input []byte) (err error) {
	var fields map[string]json.RawMessage

	err = json.Unmarshal(input, &fields)
	if err != nil {
		return
	}

	if fields["rti"] != nil {
		err = json.Unmarshal(fields["rti"], &node.Rti)
		if err != nil {
			return
		}
	}

	if fields["prti"] != nil {
		err = json.Unmarshal(fields["prti"], &node.Prti)
		if err != nil {
			return
		}
	}

	if fields["rowmarkId"] != nil {
		err = json.Unmarshal(fields["rowmarkId"], &node.RowmarkId)
		if err != nil {
			return
		}
	}

	if fields["markType"] != nil {
		err = json.Unmarshal(fields["markType"], &node.MarkType)
		if err != nil {
			return
		}
	}

	if fields["noWait"] != nil {
		err = json.Unmarshal(fields["noWait"], &node.NoWait)
		if err != nil {
			return
		}
	}

	if fields["isParent"] != nil {
		err = json.Unmarshal(fields["isParent"], &node.IsParent)
		if err != nil {
			return
		}
	}

	return
}
