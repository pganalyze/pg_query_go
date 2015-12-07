// Auto-generated - DO NOT EDIT

package pg_query

func (node ResultPath) Fingerprint(ctx FingerprintContext) {
	ctx.WriteString("RESULTPATH")

	for _, subNode := range node.Quals {
		subNode.Fingerprint(ctx)
	}
}
