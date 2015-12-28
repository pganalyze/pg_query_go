// Auto-generated - DO NOT EDIT

package pg_query

func (node FuncWithArgs) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("FuncWithArgs")
	if len(node.Funcargs.Items) > 0 {
		ctx.WriteString("funcargs")
		node.Funcargs.Fingerprint(ctx, "Funcargs")
	}

	if len(node.Funcname.Items) > 0 {
		ctx.WriteString("funcname")
		node.Funcname.Fingerprint(ctx, "Funcname")
	}
}
