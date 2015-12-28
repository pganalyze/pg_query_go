// Auto-generated - DO NOT EDIT

package pg_query

func (node ExecuteStmt) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("ExecuteStmt")
	// Intentionally ignoring node.Name for fingerprinting

	if len(node.Params.Items) > 0 {
		ctx.WriteString("params")
		node.Params.Fingerprint(ctx, "Params")
	}
}
