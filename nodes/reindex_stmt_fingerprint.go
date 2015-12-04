// Auto-generated - DO NOT EDIT

package pg_query

import (
	"io"
	"strconv"
)

func (node ReindexStmt) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "REINDEXSTMT")
	io.WriteString(ctx.hash, strconv.FormatBool(node.DoSystem))
	io.WriteString(ctx.hash, strconv.FormatBool(node.DoUser))
	io.WriteString(ctx.hash, strconv.Itoa(int(node.Kind)))

	if node.Name != nil {
		io.WriteString(ctx.hash, *node.Name)
	}

	if node.Relation != nil {
		node.Relation.Fingerprint(ctx)
	}
}
