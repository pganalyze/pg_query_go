// Auto-generated - DO NOT EDIT

package pg_query

import "sort"

func (node InsertStmt) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("InsertStmt")
	var colsFingerprints FingerprintSubContextSlice

	for _, subNode := range node.Cols {
		subCtx := FingerprintSubContext{}
		subNode.Fingerprint(&subCtx, "Cols")
		colsFingerprints.AddIfUnique(subCtx)
	}

	sort.Sort(colsFingerprints)

	for _, fingerprint := range colsFingerprints {
		for _, part := range fingerprint.parts {
			ctx.WriteString(part)
		}
	}

	if node.Relation != nil {
		node.Relation.Fingerprint(ctx, "Relation")
	}

	for _, subNode := range node.ReturningList {
		subNode.Fingerprint(ctx, "ReturningList")
	}

	if node.SelectStmt != nil {
		node.SelectStmt.Fingerprint(ctx, "SelectStmt")
	}

	if node.WithClause != nil {
		node.WithClause.Fingerprint(ctx, "WithClause")
	}
}
