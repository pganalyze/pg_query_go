// Auto-generated - DO NOT EDIT

package pg_query

import "io"

func (node SortBy) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "SortBy")
	if node.Node != nil {
		node.Node.Fingerprint(ctx)
	}

	for _, subNode := range node.UseOp {
		subNode.Fingerprint(ctx)
	}
}
