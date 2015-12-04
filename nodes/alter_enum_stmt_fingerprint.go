// Auto-generated - DO NOT EDIT

package pg_query

import (
	"io"
	"strconv"
)

func (node AlterEnumStmt) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "ALTERENUMSTMT")

	if node.NewVal != nil {
		io.WriteString(ctx.hash, *node.NewVal)
	}

	io.WriteString(ctx.hash, strconv.FormatBool(node.NewValIsAfter))

	if node.NewValNeighbor != nil {
		io.WriteString(ctx.hash, *node.NewValNeighbor)
	}

	io.WriteString(ctx.hash, strconv.FormatBool(node.SkipIfExists))

	for _, subNode := range node.TypeName {
		subNode.Fingerprint(ctx)
	}
}
