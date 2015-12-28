// Auto-generated - DO NOT EDIT

package pg_query

func (node CreateRangeStmt) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("CreateRangeStmt")
	if len(node.Params.Items) > 0 {
		ctx.WriteString("params")
		node.Params.Fingerprint(ctx, "Params")
	}

	if len(node.TypeName.Items) > 0 {
		ctx.WriteString("typeName")
		node.TypeName.Fingerprint(ctx, "TypeName")
	}
}
