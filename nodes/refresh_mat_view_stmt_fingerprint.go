// Auto-generated - DO NOT EDIT

package pg_query

import (
	"io"
	"strconv"
)

func (node RefreshMatViewStmt) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "REFRESHMATVIEWSTMT")
	io.WriteString(ctx.hash, strconv.FormatBool(node.Concurrent))

	if node.Relation != nil {
		node.Relation.Fingerprint(ctx)
	}

	io.WriteString(ctx.hash, strconv.FormatBool(node.SkipData))
}
