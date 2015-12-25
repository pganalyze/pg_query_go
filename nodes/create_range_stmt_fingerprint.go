// Auto-generated - DO NOT EDIT

package pg_query

func (node CreateRangeStmt) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("CreateRangeStmt")
	node.Params.Fingerprint(ctx, "Params")
	node.TypeName.Fingerprint(ctx, "TypeName")
}
