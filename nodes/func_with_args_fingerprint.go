// Auto-generated - DO NOT EDIT

package pg_query

func (node FuncWithArgs) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("FuncWithArgs")

	for _, subNode := range node.Funcargs {
		subNode.Fingerprint(ctx, "Funcargs")
	}

	for _, subNode := range node.Funcname {
		subNode.Fingerprint(ctx, "Funcname")
	}
}
