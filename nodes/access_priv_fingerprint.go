// Auto-generated - DO NOT EDIT

package pg_query

func (node AccessPriv) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("AccessPriv")
	node.Cols.Fingerprint(ctx, "Cols")

	if node.PrivName != nil {
		ctx.WriteString(*node.PrivName)
	}
}
