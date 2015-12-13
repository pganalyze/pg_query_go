// Auto-generated - DO NOT EDIT

package pg_query

func (node AlternativeSubPlan) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("AlternativeSubPlan")

	for _, subNode := range node.Subplans {
		subNode.Fingerprint(ctx, "Subplans")
	}

	if node.Xpr != nil {
		node.Xpr.Fingerprint(ctx, "Xpr")
	}
}
