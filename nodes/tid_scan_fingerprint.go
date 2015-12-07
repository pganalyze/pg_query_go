// Auto-generated - DO NOT EDIT

package pg_query

func (node TidScan) Fingerprint(ctx FingerprintContext) {
	ctx.WriteString("TIDSCAN")

	for _, subNode := range node.Tidquals {
		subNode.Fingerprint(ctx)
	}
}
