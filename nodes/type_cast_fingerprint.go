// Auto-generated - DO NOT EDIT

package pg_query

func (node TypeCast) Fingerprint(ctx FingerprintContext) {
	ctx.WriteString("TypeCast")

	if node.Arg != nil {
		node.Arg.Fingerprint(ctx)
	}

	// Intentionally ignoring node.Location for fingerprinting

	if node.TypeName != nil {
		node.TypeName.Fingerprint(ctx)
	}
}
