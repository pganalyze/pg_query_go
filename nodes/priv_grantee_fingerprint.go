// Auto-generated - DO NOT EDIT

package pg_query

func (node PrivGrantee) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("PrivGrantee")

	if node.Rolname != nil {
		ctx.WriteString(*node.Rolname)
	}
}
