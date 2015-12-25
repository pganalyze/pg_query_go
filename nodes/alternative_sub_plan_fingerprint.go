// Auto-generated - DO NOT EDIT

package pg_query

func (node AlternativeSubPlan) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("AlternativeSubPlan")
	node.Subplans.Fingerprint(ctx, "Subplans")

	if node.Xpr != nil {
		node.Xpr.Fingerprint(ctx, "Xpr")
	}
}
