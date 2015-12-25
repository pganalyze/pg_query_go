// Auto-generated - DO NOT EDIT

package pg_query

func (node FuncWithArgs) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("FuncWithArgs")
	node.Funcargs.Fingerprint(ctx, "Funcargs")
	node.Funcname.Fingerprint(ctx, "Funcname")
}
