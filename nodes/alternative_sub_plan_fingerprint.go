// Auto-generated - DO NOT EDIT

package pg_query

func (node AlternativeSubPlan) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("AlternativeSubPlan")
	if len(node.Subplans.Items) > 0 {
		ctx.WriteString("subplans")
		node.Subplans.Fingerprint(ctx, "Subplans")
	}

	if node.Xpr != nil {
		ctx.WriteString("xpr")
		node.Xpr.Fingerprint(ctx, "Xpr")
	}
}
