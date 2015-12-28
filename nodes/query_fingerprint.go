// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node Query) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("Query")

	if node.CanSetTag {
		ctx.WriteString("canSetTag")
		ctx.WriteString(strconv.FormatBool(node.CanSetTag))
	}

	if int(node.CommandType) != 0 {
		ctx.WriteString("commandType")
		ctx.WriteString(strconv.Itoa(int(node.CommandType)))
	}

	if len(node.ConstraintDeps.Items) > 0 {
		ctx.WriteString("constraintDeps")
		node.ConstraintDeps.Fingerprint(ctx, "ConstraintDeps")
	}

	if len(node.CteList.Items) > 0 {
		ctx.WriteString("cteList")
		node.CteList.Fingerprint(ctx, "CteList")
	}

	if len(node.DistinctClause.Items) > 0 {
		ctx.WriteString("distinctClause")
		node.DistinctClause.Fingerprint(ctx, "DistinctClause")
	}

	if len(node.GroupClause.Items) > 0 {
		ctx.WriteString("groupClause")
		node.GroupClause.Fingerprint(ctx, "GroupClause")
	}

	if node.HasAggs {
		ctx.WriteString("hasAggs")
		ctx.WriteString(strconv.FormatBool(node.HasAggs))
	}

	if node.HasDistinctOn {
		ctx.WriteString("hasDistinctOn")
		ctx.WriteString(strconv.FormatBool(node.HasDistinctOn))
	}

	if node.HasForUpdate {
		ctx.WriteString("hasForUpdate")
		ctx.WriteString(strconv.FormatBool(node.HasForUpdate))
	}

	if node.HasModifyingCte {
		ctx.WriteString("hasModifyingCTE")
		ctx.WriteString(strconv.FormatBool(node.HasModifyingCte))
	}

	if node.HasRecursive {
		ctx.WriteString("hasRecursive")
		ctx.WriteString(strconv.FormatBool(node.HasRecursive))
	}

	if node.HasSubLinks {
		ctx.WriteString("hasSubLinks")
		ctx.WriteString(strconv.FormatBool(node.HasSubLinks))
	}

	if node.HasWindowFuncs {
		ctx.WriteString("hasWindowFuncs")
		ctx.WriteString(strconv.FormatBool(node.HasWindowFuncs))
	}

	if node.HavingQual != nil {
		ctx.WriteString("havingQual")
		node.HavingQual.Fingerprint(ctx, "HavingQual")
	}

	if node.Jointree != nil {
		ctx.WriteString("jointree")
		node.Jointree.Fingerprint(ctx, "Jointree")
	}

	if node.LimitCount != nil {
		ctx.WriteString("limitCount")
		node.LimitCount.Fingerprint(ctx, "LimitCount")
	}

	if node.LimitOffset != nil {
		ctx.WriteString("limitOffset")
		node.LimitOffset.Fingerprint(ctx, "LimitOffset")
	}

	if node.QueryId != 0 {
		ctx.WriteString("queryId")
		ctx.WriteString(strconv.Itoa(int(node.QueryId)))
	}

	if int(node.QuerySource) != 0 {
		ctx.WriteString("querySource")
		ctx.WriteString(strconv.Itoa(int(node.QuerySource)))
	}

	if node.ResultRelation != 0 {
		ctx.WriteString("resultRelation")
		ctx.WriteString(strconv.Itoa(int(node.ResultRelation)))
	}

	if len(node.ReturningList.Items) > 0 {
		ctx.WriteString("returningList")
		node.ReturningList.Fingerprint(ctx, "ReturningList")
	}

	if len(node.RowMarks.Items) > 0 {
		ctx.WriteString("rowMarks")
		node.RowMarks.Fingerprint(ctx, "RowMarks")
	}

	if len(node.Rtable.Items) > 0 {
		ctx.WriteString("rtable")
		node.Rtable.Fingerprint(ctx, "Rtable")
	}

	if node.SetOperations != nil {
		ctx.WriteString("setOperations")
		node.SetOperations.Fingerprint(ctx, "SetOperations")
	}

	if len(node.SortClause.Items) > 0 {
		ctx.WriteString("sortClause")
		node.SortClause.Fingerprint(ctx, "SortClause")
	}

	if len(node.TargetList.Items) > 0 {
		ctx.WriteString("targetList")
		node.TargetList.Fingerprint(ctx, "TargetList")
	}

	if node.UtilityStmt != nil {
		ctx.WriteString("utilityStmt")
		node.UtilityStmt.Fingerprint(ctx, "UtilityStmt")
	}

	if len(node.WindowClause.Items) > 0 {
		ctx.WriteString("windowClause")
		node.WindowClause.Fingerprint(ctx, "WindowClause")
	}

	if len(node.WithCheckOptions.Items) > 0 {
		ctx.WriteString("withCheckOptions")
		node.WithCheckOptions.Fingerprint(ctx, "WithCheckOptions")
	}
}
