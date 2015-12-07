// Auto-generated - DO NOT EDIT

package pg_query

func (node Append) Fingerprint(ctx FingerprintContext) {
	ctx.WriteString("APPEND")

	for _, subNode := range node.Appendplans {
		subNode.Fingerprint(ctx)
	}
}
