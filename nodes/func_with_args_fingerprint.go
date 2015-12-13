// Auto-generated - DO NOT EDIT

package pg_query

func (node FuncWithArgs) Fingerprint(ctx FingerprintContext) {
	ctx.WriteString("FuncWithArgs")

	for _, subNode := range node.Funcargs {
		subNode.Fingerprint(ctx)
	}

	for _, subNode := range node.Funcname {
		subNode.Fingerprint(ctx)
	}
}
