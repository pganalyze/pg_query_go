// Auto-generated - DO NOT EDIT

package pg_query

func (node CreateRangeStmt) Fingerprint(ctx FingerprintContext) {
	ctx.WriteString("CREATERANGESTMT")

	for _, subNode := range node.Params {
		subNode.Fingerprint(ctx)
	}

	for _, subNode := range node.TypeName {
		subNode.Fingerprint(ctx)
	}
}
