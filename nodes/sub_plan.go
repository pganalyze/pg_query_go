// Auto-generated from postgres/src/include/nodes/primnodes.h - DO NOT EDIT

package pg_query

import "encoding/json"

type SubPlan struct {
	Xpr Expr `json:"xpr"`

	/* Fields copied from original SubLink: */
	SubLinkType SubLinkType `json:"subLinkType"` /* see above */

	/* The combining operators, transformed to an executable expression: */
	Testexpr Node   `json:"testexpr"` /* OpExpr or RowCompareExpr expression tree */
	ParamIds []Node `json:"paramIds"` /* IDs of Params embedded in the above */

	/* Identification of the Plan tree to use: */
	PlanId int `json:"plan_id"` /* Index (from 1) in PlannedStmt.subplans */

	/* Identification of the SubPlan for EXPLAIN and debugging purposes: */
	PlanName *string `json:"plan_name"` /* A name assigned during planning */

	/* Extra data useful for determining subplan's output type: */
	FirstColType      Oid   `json:"firstColType"`      /* Type of first column of subplan result */
	FirstColTypmod    int32 `json:"firstColTypmod"`    /* Typmod of first column of subplan result */
	FirstColCollation Oid   `json:"firstColCollation"` /* Collation of first column of
	 * subplan result */

	/* Information about execution strategy: */
	UseHashTable bool `json:"useHashTable"` /* TRUE to store subselect output in a hash
	 * table (implies we are doing "IN") */

	UnknownEqFalse bool `json:"unknownEqFalse"` /* TRUE if it's okay to return FALSE when the
	 * spec result is UNKNOWN; this allows much
	 * simpler handling of null values */

	/* Information for passing params into and out of the subselect: */

	/* setParam and parParam are lists of integers (param IDs) */
	SetParam []Node `json:"setParam"` /* initplan subqueries have to set these
	 * Params for parent plan */

	ParParam []Node `json:"parParam"` /* indices of input Params from parent plan */
	Args     []Node `json:"args"`     /* exprs to pass as parParam values */

	/* Estimated execution costs: */
	StartupCost Cost `json:"startup_cost"`  /* one-time setup cost */
	PerCallCost Cost `json:"per_call_cost"` /* cost for each subplan evaluation */
}

func (node SubPlan) MarshalJSON() ([]byte, error) {
	type SubPlanMarshalAlias SubPlan
	return json.Marshal(map[string]interface{}{
		"SUBPLAN": (*SubPlanMarshalAlias)(&node),
	})
}

func (node *SubPlan) UnmarshalJSON(input []byte) (err error) {
	var fields map[string]json.RawMessage

	err = json.Unmarshal(input, &fields)
	if err != nil {
		return
	}

	if fields["xpr"] != nil {
		err = json.Unmarshal(fields["xpr"], &node.Xpr)
		if err != nil {
			return
		}
	}

	if fields["subLinkType"] != nil {
		err = json.Unmarshal(fields["subLinkType"], &node.SubLinkType)
		if err != nil {
			return
		}
	}

	if fields["testexpr"] != nil {
		node.Testexpr, err = UnmarshalNodeJSON(fields["testexpr"])
		if err != nil {
			return
		}
	}

	if fields["paramIds"] != nil {
		node.ParamIds, err = UnmarshalNodeArrayJSON(fields["paramIds"])
		if err != nil {
			return
		}
	}

	if fields["plan_id"] != nil {
		err = json.Unmarshal(fields["plan_id"], &node.PlanId)
		if err != nil {
			return
		}
	}

	if fields["plan_name"] != nil {
		err = json.Unmarshal(fields["plan_name"], &node.PlanName)
		if err != nil {
			return
		}
	}

	if fields["firstColType"] != nil {
		err = json.Unmarshal(fields["firstColType"], &node.FirstColType)
		if err != nil {
			return
		}
	}

	if fields["firstColTypmod"] != nil {
		err = json.Unmarshal(fields["firstColTypmod"], &node.FirstColTypmod)
		if err != nil {
			return
		}
	}

	if fields["firstColCollation"] != nil {
		err = json.Unmarshal(fields["firstColCollation"], &node.FirstColCollation)
		if err != nil {
			return
		}
	}

	if fields["useHashTable"] != nil {
		err = json.Unmarshal(fields["useHashTable"], &node.UseHashTable)
		if err != nil {
			return
		}
	}

	if fields["unknownEqFalse"] != nil {
		err = json.Unmarshal(fields["unknownEqFalse"], &node.UnknownEqFalse)
		if err != nil {
			return
		}
	}

	if fields["setParam"] != nil {
		node.SetParam, err = UnmarshalNodeArrayJSON(fields["setParam"])
		if err != nil {
			return
		}
	}

	if fields["parParam"] != nil {
		node.ParParam, err = UnmarshalNodeArrayJSON(fields["parParam"])
		if err != nil {
			return
		}
	}

	if fields["args"] != nil {
		node.Args, err = UnmarshalNodeArrayJSON(fields["args"])
		if err != nil {
			return
		}
	}

	if fields["startup_cost"] != nil {
		err = json.Unmarshal(fields["startup_cost"], &node.StartupCost)
		if err != nil {
			return
		}
	}

	if fields["per_call_cost"] != nil {
		err = json.Unmarshal(fields["per_call_cost"], &node.PerCallCost)
		if err != nil {
			return
		}
	}

	return
}
