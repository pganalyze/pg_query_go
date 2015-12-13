// Auto-generated - DO NOT EDIT

package pg_query

func (node DoStmt) Fingerprint(ctx FingerprintContext) {
	ctx.WriteString("DoStmt")

	for _, subNode := range node.Args {
		subNode.Fingerprint(ctx)
	}
}
