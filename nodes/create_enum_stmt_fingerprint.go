// Auto-generated - DO NOT EDIT

package pg_query

func (node CreateEnumStmt) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("CreateEnumStmt")
	node.TypeName.Fingerprint(ctx, "TypeName")
	node.Vals.Fingerprint(ctx, "Vals")
}
