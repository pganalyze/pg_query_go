// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node SelectStmt) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("SelectStmt")
	ctx.WriteString(strconv.FormatBool(node.All))
	node.DistinctClause.Fingerprint(ctx, "DistinctClause")
	node.FromClause.Fingerprint(ctx, "FromClause")
	node.GroupClause.Fingerprint(ctx, "GroupClause")

	if node.HavingClause != nil {
		node.HavingClause.Fingerprint(ctx, "HavingClause")
	}

	if node.IntoClause != nil {
		node.IntoClause.Fingerprint(ctx, "IntoClause")
	}

	if node.Larg != nil {
		node.Larg.Fingerprint(ctx, "Larg")
	}

	if node.LimitCount != nil {
		node.LimitCount.Fingerprint(ctx, "LimitCount")
	}

	if node.LimitOffset != nil {
		node.LimitOffset.Fingerprint(ctx, "LimitOffset")
	}

	node.LockingClause.Fingerprint(ctx, "LockingClause")
	ctx.WriteString(strconv.Itoa(int(node.Op)))

	if node.Rarg != nil {
		node.Rarg.Fingerprint(ctx, "Rarg")
	}

	node.SortClause.Fingerprint(ctx, "SortClause")
	node.TargetList.Fingerprint(ctx, "TargetList")

	for _, nodeList := range node.ValuesLists {
		for _, subNode := range nodeList {
			subNode.Fingerprint(ctx, "ValuesLists")
		}
	}

	if node.WhereClause != nil {
		node.WhereClause.Fingerprint(ctx, "WhereClause")
	}

	node.WindowClause.Fingerprint(ctx, "WindowClause")

	if node.WithClause != nil {
		node.WithClause.Fingerprint(ctx, "WithClause")
	}
}
