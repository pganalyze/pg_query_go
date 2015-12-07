// Auto-generated - DO NOT EDIT

package pg_query

func (node MaterialPath) Fingerprint(ctx FingerprintContext) {
	ctx.WriteString("MATERIALPATH")

	if node.Subpath != nil {
		node.Subpath.Fingerprint(ctx)
	}
}
