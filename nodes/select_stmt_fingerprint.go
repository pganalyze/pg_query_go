// Auto-generated - DO NOT EDIT

package pg_query

import (
	"sort"
	"strconv"
)

func (node SelectStmt) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("SelectStmt")
	ctx.WriteString(strconv.FormatBool(node.All))

	for _, subNode := range node.DistinctClause {
		subNode.Fingerprint(ctx, "DistinctClause")
	}

	for _, subNode := range node.FromClause {
		subNode.Fingerprint(ctx, "FromClause")
	}

	for _, subNode := range node.GroupClause {
		subNode.Fingerprint(ctx, "GroupClause")
	}

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

	for _, subNode := range node.LockingClause {
		subNode.Fingerprint(ctx, "LockingClause")
	}

	ctx.WriteString(strconv.Itoa(int(node.Op)))

	if node.Rarg != nil {
		node.Rarg.Fingerprint(ctx, "Rarg")
	}

	for _, subNode := range node.SortClause {
		subNode.Fingerprint(ctx, "SortClause")
	}

	var targetListFingerprints FingerprintSubContextSlice

	for _, subNode := range node.TargetList {
		subCtx := FingerprintSubContext{}
		subNode.Fingerprint(&subCtx, "TargetList")
		targetListFingerprints.AddIfUnique(subCtx)
	}

	sort.Sort(targetListFingerprints)

	for _, fingerprint := range targetListFingerprints {
		for _, part := range fingerprint.parts {
			ctx.WriteString(part)
		}
	}

	for _, nodeList := range node.ValuesLists {
		for _, subNode := range nodeList {
			subNode.Fingerprint(ctx, "ValuesLists")
		}
	}

	if node.WhereClause != nil {
		node.WhereClause.Fingerprint(ctx, "WhereClause")
	}

	for _, subNode := range node.WindowClause {
		subNode.Fingerprint(ctx, "WindowClause")
	}

	if node.WithClause != nil {
		node.WithClause.Fingerprint(ctx, "WithClause")
	}
}
