// Auto-generated from postgres/src/include/nodes/plannodes.h - DO NOT EDIT

package pg_query

import "encoding/json"

/* ----------------
 *	 ModifyTable node -
 *		Apply rows produced by subplan(s) to result table(s),
 *		by inserting, updating, or deleting.
 *
 * Note that rowMarks and epqParam are presumed to be valid for all the
 * subplan(s); they can't contain any info that varies across subplans.
 * ----------------
 */
type ModifyTable struct {
	Plan                 Plan    `json:"plan"`
	Operation            CmdType `json:"operation"`            /* INSERT, UPDATE, or DELETE */
	CanSetTag            bool    `json:"canSetTag"`            /* do we set the command tag/es_processed? */
	ResultRelations      []Node  `json:"resultRelations"`      /* integer list of RT indexes */
	ResultRelIndex       int     `json:"resultRelIndex"`       /* index of first resultRel in plan's list */
	Plans                []Node  `json:"plans"`                /* plan(s) producing source data */
	WithCheckOptionLists []Node  `json:"withCheckOptionLists"` /* per-target-table WCO lists */
	ReturningLists       []Node  `json:"returningLists"`       /* per-target-table RETURNING tlists */
	FdwPrivLists         []Node  `json:"fdwPrivLists"`         /* per-target-table FDW private data lists */
	RowMarks             []Node  `json:"rowMarks"`             /* PlanRowMarks (non-locking only) */
	EpqParam             int     `json:"epqParam"`             /* ID of Param for EvalPlanQual re-eval */
}

func (node ModifyTable) MarshalJSON() ([]byte, error) {
	type ModifyTableMarshalAlias ModifyTable
	return json.Marshal(map[string]interface{}{
		"MODIFYTABLE": (*ModifyTableMarshalAlias)(&node),
	})
}

func (node *ModifyTable) UnmarshalJSON(input []byte) (err error) {
	var fields map[string]json.RawMessage

	err = json.Unmarshal(input, &fields)
	if err != nil {
		return
	}

	if fields["plan"] != nil {
		err = json.Unmarshal(fields["plan"], &node.Plan)
		if err != nil {
			return
		}
	}

	if fields["operation"] != nil {
		err = json.Unmarshal(fields["operation"], &node.Operation)
		if err != nil {
			return
		}
	}

	if fields["canSetTag"] != nil {
		err = json.Unmarshal(fields["canSetTag"], &node.CanSetTag)
		if err != nil {
			return
		}
	}

	if fields["resultRelations"] != nil {
		node.ResultRelations, err = UnmarshalNodeArrayJSON(fields["resultRelations"])
		if err != nil {
			return
		}
	}

	if fields["resultRelIndex"] != nil {
		err = json.Unmarshal(fields["resultRelIndex"], &node.ResultRelIndex)
		if err != nil {
			return
		}
	}

	if fields["plans"] != nil {
		node.Plans, err = UnmarshalNodeArrayJSON(fields["plans"])
		if err != nil {
			return
		}
	}

	if fields["withCheckOptionLists"] != nil {
		node.WithCheckOptionLists, err = UnmarshalNodeArrayJSON(fields["withCheckOptionLists"])
		if err != nil {
			return
		}
	}

	if fields["returningLists"] != nil {
		node.ReturningLists, err = UnmarshalNodeArrayJSON(fields["returningLists"])
		if err != nil {
			return
		}
	}

	if fields["fdwPrivLists"] != nil {
		node.FdwPrivLists, err = UnmarshalNodeArrayJSON(fields["fdwPrivLists"])
		if err != nil {
			return
		}
	}

	if fields["rowMarks"] != nil {
		node.RowMarks, err = UnmarshalNodeArrayJSON(fields["rowMarks"])
		if err != nil {
			return
		}
	}

	if fields["epqParam"] != nil {
		err = json.Unmarshal(fields["epqParam"], &node.EpqParam)
		if err != nil {
			return
		}
	}

	return
}
