// Auto-generated - DO NOT EDIT

package pg_query

import (
	"io"
	"strconv"
)

func (node DropdbStmt) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "DROPDBSTMT")

	if node.Dbname != nil {
		io.WriteString(ctx.hash, *node.Dbname)
	}

	io.WriteString(ctx.hash, strconv.FormatBool(node.MissingOk))
}
