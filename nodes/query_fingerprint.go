// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node Query) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("Query")
	ctx.WriteString(strconv.FormatBool(node.CanSetTag))
	ctx.WriteString(strconv.Itoa(int(node.CommandType)))

	for _, subNode := range node.ConstraintDeps {
		subNode.Fingerprint(ctx, "ConstraintDeps")
	}

	for _, subNode := range node.CteList {
		subNode.Fingerprint(ctx, "CteList")
	}

	for _, subNode := range node.DistinctClause {
		subNode.Fingerprint(ctx, "DistinctClause")
	}

	for _, subNode := range node.GroupClause {
		subNode.Fingerprint(ctx, "GroupClause")
	}

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

	for _, subNode := range node.ReturningList {
		subNode.Fingerprint(ctx, "ReturningList")
	}

	for _, subNode := range node.RowMarks {
		subNode.Fingerprint(ctx, "RowMarks")
	}

	for _, subNode := range node.Rtable {
		subNode.Fingerprint(ctx, "Rtable")
	}

	if node.SetOperations != nil {
		node.SetOperations.Fingerprint(ctx, "SetOperations")
	}

	for _, subNode := range node.SortClause {
		subNode.Fingerprint(ctx, "SortClause")
	}

	for _, subNode := range node.TargetList {
		subNode.Fingerprint(ctx, "TargetList")
	}

	if node.UtilityStmt != nil {
		node.UtilityStmt.Fingerprint(ctx, "UtilityStmt")
	}

	for _, subNode := range node.WindowClause {
		subNode.Fingerprint(ctx, "WindowClause")
	}

	for _, subNode := range node.WithCheckOptions {
		subNode.Fingerprint(ctx, "WithCheckOptions")
	}
}
