// Auto-generated - DO NOT EDIT

package pg_query

func (node NestLoopParam) Fingerprint(ctx FingerprintContext) {
	ctx.WriteString("NESTLOOPPARAM")

	if node.Paramval != nil {
		node.Paramval.Fingerprint(ctx)
	}
}
