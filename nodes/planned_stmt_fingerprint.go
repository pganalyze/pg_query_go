// Auto-generated - DO NOT EDIT

package pg_query

import "io"

func (node PlannedStmt) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "PlannedStmt")

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

	if node.UtilityStmt != nil {
		node.UtilityStmt.Fingerprint(ctx)
	}
}
