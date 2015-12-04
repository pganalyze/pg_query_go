// Auto-generated - DO NOT EDIT

package pg_query

import (
	"io"
	"strconv"
)

func (node WindowClause) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "WINDOWCLAUSE")
	io.WriteString(ctx.hash, strconv.FormatBool(node.CopiedOrder))

	if node.EndOffset != nil {
		node.EndOffset.Fingerprint(ctx)
	}

	if node.Name != nil {
		io.WriteString(ctx.hash, *node.Name)
	}

	for _, subNode := range node.OrderClause {
		subNode.Fingerprint(ctx)
	}

	for _, subNode := range node.PartitionClause {
		subNode.Fingerprint(ctx)
	}

	if node.Refname != nil {
		io.WriteString(ctx.hash, *node.Refname)
	}

	if node.StartOffset != nil {
		node.StartOffset.Fingerprint(ctx)
	}
}
