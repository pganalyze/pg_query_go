// Auto-generated - DO NOT EDIT

package pg_query

import "io"

func (node WindowClause) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "WindowClause")
	if node.EndOffset != nil {
		node.EndOffset.Fingerprint(ctx)
	}

	for _, subNode := range node.OrderClause {
		subNode.Fingerprint(ctx)
	}

	for _, subNode := range node.PartitionClause {
		subNode.Fingerprint(ctx)
	}

	if node.StartOffset != nil {
		node.StartOffset.Fingerprint(ctx)
	}
}
