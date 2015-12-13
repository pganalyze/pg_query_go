// Auto-generated - DO NOT EDIT

package pg_query

func (node FromExpr) Fingerprint(ctx FingerprintContext) {
	ctx.WriteString("FromExpr")

	for _, subNode := range node.Fromlist {
		subNode.Fingerprint(ctx)
	}

	if node.Quals != nil {
		node.Quals.Fingerprint(ctx)
	}
}
