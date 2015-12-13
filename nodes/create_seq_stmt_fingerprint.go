// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node CreateSeqStmt) Fingerprint(ctx FingerprintContext) {
	ctx.WriteString("CreateSeqStmt")

	for _, subNode := range node.Options {
		subNode.Fingerprint(ctx)
	}

	ctx.WriteString(strconv.Itoa(int(node.OwnerId)))

	if node.Sequence != nil {
		node.Sequence.Fingerprint(ctx)
	}
}
