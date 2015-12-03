// Auto-generated - DO NOT EDIT

package pg_query

import "io"

func (node CreateTableAsStmt) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "CreateTableAsStmt")
	if node.Into != nil {
		node.Into.Fingerprint(ctx)
	}

	if node.Query != nil {
		node.Query.Fingerprint(ctx)
	}
}
