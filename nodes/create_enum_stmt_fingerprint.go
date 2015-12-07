// Auto-generated - DO NOT EDIT

package pg_query

func (node CreateEnumStmt) Fingerprint(ctx FingerprintContext) {
	ctx.WriteString("CREATEENUMSTMT")

	for _, subNode := range node.TypeName {
		subNode.Fingerprint(ctx)
	}

	for _, subNode := range node.Vals {
		subNode.Fingerprint(ctx)
	}
}
