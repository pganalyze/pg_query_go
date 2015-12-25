// Auto-generated - DO NOT EDIT

package pg_query

func (node AlterFunctionStmt) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("AlterFunctionStmt")
	node.Actions.Fingerprint(ctx, "Actions")

	if node.Func != nil {
		node.Func.Fingerprint(ctx, "Func")
	}
}
