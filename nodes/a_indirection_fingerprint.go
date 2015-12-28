// Auto-generated - DO NOT EDIT

package pg_query

func (node A_Indirection) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("A_Indirection")

	if node.Arg != nil {
		ctx.WriteString("arg")
		node.Arg.Fingerprint(ctx, "Arg")
	}

	if len(node.Indirection.Items) > 0 {
		ctx.WriteString("indirection")
		node.Indirection.Fingerprint(ctx, "Indirection")
	}
}
