// Auto-generated - DO NOT EDIT

package pg_query

import (
	"io"
	"strconv"
)

func (node CreateExtensionStmt) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "CREATEEXTENSIONSTMT")

	if node.Extname != nil {
		io.WriteString(ctx.hash, *node.Extname)
	}

	io.WriteString(ctx.hash, strconv.FormatBool(node.IfNotExists))

	for _, subNode := range node.Options {
		subNode.Fingerprint(ctx)
	}
}
