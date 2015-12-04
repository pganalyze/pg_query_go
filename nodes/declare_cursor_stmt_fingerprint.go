// Auto-generated - DO NOT EDIT

package pg_query

import "io"

func (node DeclareCursorStmt) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "DECLARECURSOR")

	if node.Portalname != nil {
		io.WriteString(ctx.hash, *node.Portalname)
	}

	if node.Query != nil {
		node.Query.Fingerprint(ctx)
	}
}
