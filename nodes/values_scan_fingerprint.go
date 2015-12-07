// Auto-generated - DO NOT EDIT

package pg_query

func (node ValuesScan) Fingerprint(ctx FingerprintContext) {
	ctx.WriteString("VALUESSCAN")

	for _, subNode := range node.ValuesLists {
		subNode.Fingerprint(ctx)
	}
}
