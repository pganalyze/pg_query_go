// Auto-generated - DO NOT EDIT

package pg_query

import (
	"io"
	"strconv"
)

func (node ClusterStmt) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "CLUSTERSTMT")

	if node.Indexname != nil {
		io.WriteString(ctx.hash, *node.Indexname)
	}

	if node.Relation != nil {
		node.Relation.Fingerprint(ctx)
	}

	io.WriteString(ctx.hash, strconv.FormatBool(node.Verbose))
}
