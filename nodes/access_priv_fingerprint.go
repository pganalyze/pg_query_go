// Auto-generated - DO NOT EDIT

package pg_query

func (node AccessPriv) Fingerprint(ctx FingerprintContext) {
	ctx.WriteString("AccessPriv")

	for _, subNode := range node.Cols {
		subNode.Fingerprint(ctx)
	}

	if node.PrivName != nil {
		ctx.WriteString(*node.PrivName)
	}
}
