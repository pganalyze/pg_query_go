// Auto-generated - DO NOT EDIT

package pg_query

import "io"

func (node LoadStmt) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "LOADSTMT")

	if node.Filename != nil {
		io.WriteString(ctx.hash, *node.Filename)
	}
}
