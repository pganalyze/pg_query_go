// Auto-generated - DO NOT EDIT

package pg_query

import (
	"io"
	"strconv"
)

func (node SortBy) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "SORTBY")
	// Intentionally ignoring node.Location for fingerprinting

	if node.Node != nil {
		node.Node.Fingerprint(ctx)
	}

	io.WriteString(ctx.hash, strconv.Itoa(int(node.SortbyDir)))
	io.WriteString(ctx.hash, strconv.Itoa(int(node.SortbyNulls)))

	for _, subNode := range node.UseOp {
		subNode.Fingerprint(ctx)
	}
}
