// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node PlannedStmt) Fingerprint(ctx FingerprintContext) {
	ctx.WriteString("PLANNEDSTMT")
	ctx.WriteString(strconv.FormatBool(node.CanSetTag))
	ctx.WriteString(strconv.Itoa(int(node.CommandType)))
	ctx.WriteString(strconv.FormatBool(node.HasModifyingCte))
	ctx.WriteString(strconv.FormatBool(node.HasReturning))

	for _, subNode := range node.InvalItems {
		subNode.Fingerprint(ctx)
	}

	if node.PlanTree != nil {
		node.PlanTree.Fingerprint(ctx)
	}

	for _, subNode := range node.RelationOids {
		subNode.Fingerprint(ctx)
	}

	for _, subNode := range node.ResultRelations {
		subNode.Fingerprint(ctx)
	}

	for _, subNode := range node.RowMarks {
		subNode.Fingerprint(ctx)
	}

	for _, subNode := range node.Rtable {
		subNode.Fingerprint(ctx)
	}

	for _, subNode := range node.Subplans {
		subNode.Fingerprint(ctx)
	}

	ctx.WriteString(strconv.FormatBool(node.TransientPlan))

	if node.UtilityStmt != nil {
		node.UtilityStmt.Fingerprint(ctx)
	}
}
