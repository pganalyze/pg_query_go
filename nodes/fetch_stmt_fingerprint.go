// Auto-generated - DO NOT EDIT

package pg_query

import (
	"io"
	"strconv"
)

func (node FetchStmt) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "FETCHSTMT")
	io.WriteString(ctx.hash, strconv.Itoa(int(node.Direction)))
	io.WriteString(ctx.hash, strconv.FormatBool(node.Ismove))

	if node.Portalname != nil {
		io.WriteString(ctx.hash, *node.Portalname)
	}
}
