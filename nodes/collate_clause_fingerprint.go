// Auto-generated - DO NOT EDIT

package pg_query

func (node CollateClause) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("CollateClause")

	if node.Arg != nil {
		ctx.WriteString("arg")
		node.Arg.Fingerprint(ctx, "Arg")
	}

	if len(node.Collname.Items) > 0 {
		ctx.WriteString("collname")
		node.Collname.Fingerprint(ctx, "Collname")
	}

	// Intentionally ignoring node.Location for fingerprinting
}
