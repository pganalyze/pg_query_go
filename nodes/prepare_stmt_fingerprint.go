// Auto-generated - DO NOT EDIT

package pg_query

func (node PrepareStmt) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("PrepareStmt")
	if len(node.Argtypes.Items) > 0 {
		ctx.WriteString("argtypes")
		node.Argtypes.Fingerprint(ctx, "Argtypes")
	}

	// Intentionally ignoring node.Name for fingerprinting

	if node.Query != nil {
		ctx.WriteString("query")
		node.Query.Fingerprint(ctx, "Query")
	}
}
