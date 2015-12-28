// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node SelectStmt) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("SelectStmt")

	if node.All {
		ctx.WriteString("all")
		ctx.WriteString(strconv.FormatBool(node.All))
	}

	if len(node.DistinctClause.Items) > 0 {
		ctx.WriteString("distinctClause")
		node.DistinctClause.Fingerprint(ctx, "DistinctClause")
	}

	if len(node.FromClause.Items) > 0 {
		ctx.WriteString("fromClause")
		node.FromClause.Fingerprint(ctx, "FromClause")
	}

	if len(node.GroupClause.Items) > 0 {
		ctx.WriteString("groupClause")
		node.GroupClause.Fingerprint(ctx, "GroupClause")
	}

	if node.HavingClause != nil {
		ctx.WriteString("havingClause")
		node.HavingClause.Fingerprint(ctx, "HavingClause")
	}

	if node.IntoClause != nil {
		ctx.WriteString("intoClause")
		node.IntoClause.Fingerprint(ctx, "IntoClause")
	}

	if node.Larg != nil {
		ctx.WriteString("larg")
		node.Larg.Fingerprint(ctx, "Larg")
	}

	if node.LimitCount != nil {
		ctx.WriteString("limitCount")
		node.LimitCount.Fingerprint(ctx, "LimitCount")
	}

	if node.LimitOffset != nil {
		ctx.WriteString("limitOffset")
		node.LimitOffset.Fingerprint(ctx, "LimitOffset")
	}

	if len(node.LockingClause.Items) > 0 {
		ctx.WriteString("lockingClause")
		node.LockingClause.Fingerprint(ctx, "LockingClause")
	}

	if int(node.Op) != 0 {
		ctx.WriteString("op")
		ctx.WriteString(strconv.Itoa(int(node.Op)))
	}

	if node.Rarg != nil {
		ctx.WriteString("rarg")
		node.Rarg.Fingerprint(ctx, "Rarg")
	}

	if len(node.SortClause.Items) > 0 {
		ctx.WriteString("sortClause")
		node.SortClause.Fingerprint(ctx, "SortClause")
	}

	if len(node.TargetList.Items) > 0 {
		ctx.WriteString("targetList")
		node.TargetList.Fingerprint(ctx, "TargetList")
	}

	if len(node.ValuesLists) > 0 {
		ctx.WriteString("valuesLists")
		for _, nodeList := range node.ValuesLists {
			for _, subNode := range nodeList {
				subNode.Fingerprint(ctx, "ValuesLists")
			}
		}
	}

	if node.WhereClause != nil {
		ctx.WriteString("whereClause")
		node.WhereClause.Fingerprint(ctx, "WhereClause")
	}

	if len(node.WindowClause.Items) > 0 {
		ctx.WriteString("windowClause")
		node.WindowClause.Fingerprint(ctx, "WindowClause")
	}

	if node.WithClause != nil {
		ctx.WriteString("withClause")
		node.WithClause.Fingerprint(ctx, "WithClause")
	}
}
