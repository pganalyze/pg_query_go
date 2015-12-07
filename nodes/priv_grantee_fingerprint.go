// Auto-generated - DO NOT EDIT

package pg_query

func (node PrivGrantee) Fingerprint(ctx FingerprintContext) {
	ctx.WriteString("PRIVGRANTEE")

	if node.Rolname != nil {
		ctx.WriteString(*node.Rolname)
	}
}
