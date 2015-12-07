// Auto-generated - DO NOT EDIT

package pg_query

func (node NestLoop) Fingerprint(ctx FingerprintContext) {
	ctx.WriteString("NESTLOOP")

	for _, subNode := range node.NestParams {
		subNode.Fingerprint(ctx)
	}
}
