// Auto-generated - DO NOT EDIT

package pg_query

import "io"

func (node CreateSeqStmt) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "CreateSeqStmt")

	for _, subNode := range node.Options {
		subNode.Fingerprint(ctx)
	}

	if node.Sequence != nil {
		node.Sequence.Fingerprint(ctx)
	}
}
