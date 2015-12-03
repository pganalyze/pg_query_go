// Auto-generated - DO NOT EDIT

package pg_query

import "io"

func (node Query) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "Query")

	for _, subNode := range node.ConstraintDeps {
		subNode.Fingerprint(ctx)
	}

	for _, subNode := range node.CteList {
		subNode.Fingerprint(ctx)
	}

	for _, subNode := range node.DistinctClause {
		subNode.Fingerprint(ctx)
	}

	for _, subNode := range node.GroupClause {
		subNode.Fingerprint(ctx)
	}

	if node.HavingQual != nil {
		node.HavingQual.Fingerprint(ctx)
	}

	if node.Jointree != nil {
		node.Jointree.Fingerprint(ctx)
	}

	if node.LimitCount != nil {
		node.LimitCount.Fingerprint(ctx)
	}

	if node.LimitOffset != nil {
		node.LimitOffset.Fingerprint(ctx)
	}

	for _, subNode := range node.ReturningList {
		subNode.Fingerprint(ctx)
	}

	for _, subNode := range node.RowMarks {
		subNode.Fingerprint(ctx)
	}

	for _, subNode := range node.Rtable {
		subNode.Fingerprint(ctx)
	}

	if node.SetOperations != nil {
		node.SetOperations.Fingerprint(ctx)
	}

	for _, subNode := range node.SortClause {
		subNode.Fingerprint(ctx)
	}

	for _, subNode := range node.TargetList {
		subNode.Fingerprint(ctx)
	}

	if node.UtilityStmt != nil {
		node.UtilityStmt.Fingerprint(ctx)
	}

	for _, subNode := range node.WindowClause {
		subNode.Fingerprint(ctx)
	}

	for _, subNode := range node.WithCheckOptions {
		subNode.Fingerprint(ctx)
	}
}
