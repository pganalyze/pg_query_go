// Auto-generated - DO NOT EDIT

package pg_query

func (node TypeCast) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("TypeCast")

	if node.Arg != nil {
		ctx.WriteString("arg")
		node.Arg.Fingerprint(ctx, "Arg")
	}

	// Intentionally ignoring node.Location for fingerprinting

	if node.TypeName != nil {
		ctx.WriteString("typeName")
		node.TypeName.Fingerprint(ctx, "TypeName")
	}
}
