// Auto-generated from postgres/src/include/nodes/primnodes.h - DO NOT EDIT

package pg_query

import "encoding/json"

/*
 * SubLink
 *
 * A SubLink represents a subselect appearing in an expression, and in some
 * cases also the combining operator(s) just above it.  The subLinkType
 * indicates the form of the expression represented:
 *	EXISTS_SUBLINK		EXISTS(SELECT ...)
 *	ALL_SUBLINK			(lefthand) op ALL (SELECT ...)
 *	ANY_SUBLINK			(lefthand) op ANY (SELECT ...)
 *	ROWCOMPARE_SUBLINK	(lefthand) op (SELECT ...)
 *	EXPR_SUBLINK		(SELECT with single targetlist item ...)
 *	ARRAY_SUBLINK		ARRAY(SELECT with single targetlist item ...)
 *	CTE_SUBLINK			WITH query (never actually part of an expression)
 * For ALL, ANY, and ROWCOMPARE, the lefthand is a list of expressions of the
 * same length as the subselect's targetlist.  ROWCOMPARE will *always* have
 * a list with more than one entry; if the subselect has just one target
 * then the parser will create an EXPR_SUBLINK instead (and any operator
 * above the subselect will be represented separately).  Note that both
 * ROWCOMPARE and EXPR require the subselect to deliver only one row.
 * ALL, ANY, and ROWCOMPARE require the combining operators to deliver boolean
 * results.  ALL and ANY combine the per-row results using AND and OR
 * semantics respectively.
 * ARRAY requires just one target column, and creates an array of the target
 * column's type using any number of rows resulting from the subselect.
 *
 * SubLink is classed as an Expr node, but it is not actually executable;
 * it must be replaced in the expression tree by a SubPlan node during
 * planning.
 *
 * NOTE: in the raw output of gram.y, testexpr contains just the raw form
 * of the lefthand expression (if any), and operName is the String name of
 * the combining operator.  Also, subselect is a raw parsetree.  During parse
 * analysis, the parser transforms testexpr into a complete boolean expression
 * that compares the lefthand value(s) to PARAM_SUBLINK nodes representing the
 * output columns of the subselect.  And subselect is transformed to a Query.
 * This is the representation seen in saved rules and in the rewriter.
 *
 * In EXISTS, EXPR, and ARRAY SubLinks, testexpr and operName are unused and
 * are always null.
 *
 * The CTE_SUBLINK case never occurs in actual SubLink nodes, but it is used
 * in SubPlans generated for WITH subqueries.
 */
type SubLink struct {
	Xpr         Expr        `json:"xpr"`
	SubLinkType SubLinkType `json:"subLinkType"` /* see above */
	Testexpr    Node        `json:"testexpr"`    /* outer-query test for ALL/ANY/ROWCOMPARE */
	OperName    []Node      `json:"operName"`    /* originally specified operator name */
	Subselect   Node        `json:"subselect"`   /* subselect as Query* or parsetree */
	Location    int         `json:"location"`    /* token location, or -1 if unknown */
}

func (node SubLink) MarshalJSON() ([]byte, error) {
	type SubLinkMarshalAlias SubLink
	return json.Marshal(map[string]interface{}{
		"SUBLINK": (*SubLinkMarshalAlias)(&node),
	})
}

func (node *SubLink) UnmarshalJSON(input []byte) (err error) {
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

	if fields["operName"] != nil {
		node.OperName, err = UnmarshalNodeArrayJSON(fields["operName"])
		if err != nil {
			return
		}
	}

	if fields["subselect"] != nil {
		node.Subselect, err = UnmarshalNodeJSON(fields["subselect"])
		if err != nil {
			return
		}
	}

	if fields["location"] != nil {
		err = json.Unmarshal(fields["location"], &node.Location)
		if err != nil {
			return
		}
	}

	return
}
