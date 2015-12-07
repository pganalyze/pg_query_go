// Auto-generated - DO NOT EDIT

package pg_query

func (node CreatedbStmt) Fingerprint(ctx FingerprintContext) {
	ctx.WriteString("CREATEDBSTMT")

	if node.Dbname != nil {
		ctx.WriteString(*node.Dbname)
	}

	for _, subNode := range node.Options {
		subNode.Fingerprint(ctx)
	}
}
