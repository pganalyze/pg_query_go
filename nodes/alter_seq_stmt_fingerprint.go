// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node AlterSeqStmt) Fingerprint(ctx FingerprintContext) {
	ctx.WriteString("AlterSeqStmt")
	ctx.WriteString(strconv.FormatBool(node.MissingOk))

	for _, subNode := range node.Options {
		subNode.Fingerprint(ctx)
	}

	if node.Sequence != nil {
		node.Sequence.Fingerprint(ctx)
	}
}
