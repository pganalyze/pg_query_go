// Auto-generated - DO NOT EDIT

package pg_query

func (node ForeignPath) Fingerprint(ctx FingerprintContext) {
	ctx.WriteString("FOREIGNPATH")

	for _, subNode := range node.FdwPrivate {
		subNode.Fingerprint(ctx)
	}
}
