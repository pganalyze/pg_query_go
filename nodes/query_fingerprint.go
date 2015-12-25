// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node Query) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("Query")
	ctx.WriteString(strconv.FormatBool(node.CanSetTag))
	ctx.WriteString(strconv.Itoa(int(node.CommandType)))
	node.ConstraintDeps.Fingerprint(ctx, "ConstraintDeps")
	node.CteList.Fingerprint(ctx, "CteList")
	node.DistinctClause.Fingerprint(ctx, "DistinctClause")
	node.GroupClause.Fingerprint(ctx, "GroupClause")
	ctx.WriteString(strconv.FormatBool(node.HasAggs))
	ctx.WriteString(strconv.FormatBool(node.HasDistinctOn))
	ctx.WriteString(strconv.FormatBool(node.HasForUpdate))
	ctx.WriteString(strconv.FormatBool(node.HasModifyingCte))
	ctx.WriteString(strconv.FormatBool(node.HasRecursive))
	ctx.WriteString(strconv.FormatBool(node.HasSubLinks))
	ctx.WriteString(strconv.FormatBool(node.HasWindowFuncs))

	if node.HavingQual != nil {
		node.HavingQual.Fingerprint(ctx, "HavingQual")
	}

	if node.Jointree != nil {
		node.Jointree.Fingerprint(ctx, "Jointree")
	}

	if node.LimitCount != nil {
		node.LimitCount.Fingerprint(ctx, "LimitCount")
	}

	if node.LimitOffset != nil {
		node.LimitOffset.Fingerprint(ctx, "LimitOffset")
	}

	ctx.WriteString(strconv.Itoa(int(node.QueryId)))
	ctx.WriteString(strconv.Itoa(int(node.QuerySource)))
	ctx.WriteString(strconv.Itoa(int(node.ResultRelation)))
	node.ReturningList.Fingerprint(ctx, "ReturningList")
	node.RowMarks.Fingerprint(ctx, "RowMarks")
	node.Rtable.Fingerprint(ctx, "Rtable")

	if node.SetOperations != nil {
		node.SetOperations.Fingerprint(ctx, "SetOperations")
	}

	node.SortClause.Fingerprint(ctx, "SortClause")
	node.TargetList.Fingerprint(ctx, "TargetList")

	if node.UtilityStmt != nil {
		node.UtilityStmt.Fingerprint(ctx, "UtilityStmt")
	}

	node.WindowClause.Fingerprint(ctx, "WindowClause")
	node.WithCheckOptions.Fingerprint(ctx, "WithCheckOptions")
}
