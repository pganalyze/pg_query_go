// Auto-generated - DO NOT EDIT

package pg_query

func (node AlterFunctionStmt) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("AlterFunctionStmt")
	if len(node.Actions.Items) > 0 {
		ctx.WriteString("actions")
		node.Actions.Fingerprint(ctx, "Actions")
	}

	if node.Func != nil {
		ctx.WriteString("func")
		node.Func.Fingerprint(ctx, "Func")
	}
}
