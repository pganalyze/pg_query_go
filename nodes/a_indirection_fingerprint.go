// Auto-generated - DO NOT EDIT

package pg_query

func (node A_Indirection) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("A_Indirection")

	if node.Arg != nil {
		node.Arg.Fingerprint(ctx, "Arg")
	}

	node.Indirection.Fingerprint(ctx, "Indirection")
}
