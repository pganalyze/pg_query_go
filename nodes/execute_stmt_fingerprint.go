// Auto-generated - DO NOT EDIT

package pg_query

func (node ExecuteStmt) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("ExecuteStmt")

	if node.Name != nil {
		ctx.WriteString(*node.Name)
	}

	node.Params.Fingerprint(ctx, "Params")
}
