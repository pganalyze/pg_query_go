// Auto-generated - DO NOT EDIT

package pg_query

func (node NamedArgExpr) Fingerprint(ctx FingerprintContext) {
	ctx.WriteString("NAMEDARGEXPR")

	if node.Arg != nil {
		node.Arg.Fingerprint(ctx)
	}

	// Intentionally ignoring node.Location for fingerprinting

	if node.Name != nil {
		ctx.WriteString(*node.Name)
	}
}
