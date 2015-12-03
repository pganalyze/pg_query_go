// Auto-generated - DO NOT EDIT

package pg_query

import "io"

func (node SelectStmt) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "SelectStmt")

	for _, subNode := range node.DistinctClause {
		subNode.Fingerprint(ctx)
	}

	for _, subNode := range node.FromClause {
		subNode.Fingerprint(ctx)
	}

	for _, subNode := range node.GroupClause {
		subNode.Fingerprint(ctx)
	}

	if node.HavingClause != nil {
		node.HavingClause.Fingerprint(ctx)
	}

	if node.IntoClause != nil {
		node.IntoClause.Fingerprint(ctx)
	}

	if node.Larg != nil {
		node.Larg.Fingerprint(ctx)
	}

	if node.LimitCount != nil {
		node.LimitCount.Fingerprint(ctx)
	}

	if node.LimitOffset != nil {
		node.LimitOffset.Fingerprint(ctx)
	}

	for _, subNode := range node.LockingClause {
		subNode.Fingerprint(ctx)
	}

	if node.Rarg != nil {
		node.Rarg.Fingerprint(ctx)
	}

	for _, subNode := range node.SortClause {
		subNode.Fingerprint(ctx)
	}

	for _, subNode := range node.TargetList {
		subNode.Fingerprint(ctx)
	}

	for _, nodeList := range node.ValuesLists {
		for _, subNode := range nodeList {
			subNode.Fingerprint(ctx)
		}

	}

	if node.WhereClause != nil {
		node.WhereClause.Fingerprint(ctx)
	}

	for _, subNode := range node.WindowClause {
		subNode.Fingerprint(ctx)
	}

	if node.WithClause != nil {
		node.WithClause.Fingerprint(ctx)
	}
}
