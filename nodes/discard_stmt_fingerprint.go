// Auto-generated - DO NOT EDIT

package pg_query

import (
	"io"
	"strconv"
)

func (node DiscardStmt) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "DISCARDSTMT")
	io.WriteString(ctx.hash, strconv.Itoa(int(node.Target)))
}
