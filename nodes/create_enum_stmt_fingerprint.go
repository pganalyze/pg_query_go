// Auto-generated - DO NOT EDIT

package pg_query

func (node CreateEnumStmt) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("CreateEnumStmt")
	if len(node.TypeName.Items) > 0 {
		ctx.WriteString("typeName")
		node.TypeName.Fingerprint(ctx, "TypeName")
	}

	if len(node.Vals.Items) > 0 {
		ctx.WriteString("vals")
		node.Vals.Fingerprint(ctx, "Vals")
	}
}
