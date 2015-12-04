// Auto-generated - DO NOT EDIT

package pg_query

import (
	"io"
	"strconv"
)

func (node AlterSeqStmt) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "ALTERSEQSTMT")
	io.WriteString(ctx.hash, strconv.FormatBool(node.MissingOk))

	for _, subNode := range node.Options {
		subNode.Fingerprint(ctx)
	}

	if node.Sequence != nil {
		node.Sequence.Fingerprint(ctx)
	}
}
