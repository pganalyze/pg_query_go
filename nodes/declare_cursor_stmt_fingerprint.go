// Auto-generated - DO NOT EDIT

package pg_query

import "io"

func (node DeclareCursorStmt) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "DeclareCursorStmt")
	if node.Query != nil {
		node.Query.Fingerprint(ctx)
	}
}
