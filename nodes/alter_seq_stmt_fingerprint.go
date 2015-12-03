// Auto-generated - DO NOT EDIT

package pg_query

import "io"

func (node AlterSeqStmt) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "AlterSeqStmt")

	for _, subNode := range node.Options {
		subNode.Fingerprint(ctx)
	}

	if node.Sequence != nil {
		node.Sequence.Fingerprint(ctx)
	}
}
