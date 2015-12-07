// Auto-generated - DO NOT EDIT

package pg_query

func (node AlterFunctionStmt) Fingerprint(ctx FingerprintContext) {
	ctx.WriteString("ALTERFUNCTIONSTMT")

	for _, subNode := range node.Actions {
		subNode.Fingerprint(ctx)
	}

	if node.Func != nil {
		node.Func.Fingerprint(ctx)
	}
}
