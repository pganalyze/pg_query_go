// Auto-generated - DO NOT EDIT

package pg_query

func (node TidPath) Fingerprint(ctx FingerprintContext) {
	ctx.WriteString("TIDPATH")

	for _, subNode := range node.Tidquals {
		subNode.Fingerprint(ctx)
	}
}
