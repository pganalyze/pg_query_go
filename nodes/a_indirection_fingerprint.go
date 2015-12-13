// Auto-generated - DO NOT EDIT

package pg_query

func (node A_Indirection) Fingerprint(ctx FingerprintContext) {
	ctx.WriteString("A_Indirection")

	if node.Arg != nil {
		node.Arg.Fingerprint(ctx)
	}

	for _, subNode := range node.Indirection {
		subNode.Fingerprint(ctx)
	}
}
