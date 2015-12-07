// Auto-generated - DO NOT EDIT

package pg_query

func (node CreateSeqStmt) Fingerprint(ctx FingerprintContext) {
	ctx.WriteString("CREATESEQSTMT")

	for _, subNode := range node.Options {
		subNode.Fingerprint(ctx)
	}

	if node.Sequence != nil {
		node.Sequence.Fingerprint(ctx)
	}
}
