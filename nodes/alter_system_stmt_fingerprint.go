// Auto-generated - DO NOT EDIT

package pg_query

import "io"

func (node AlterSystemStmt) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "AlterSystemStmt")
	if node.Setstmt != nil {
		node.Setstmt.Fingerprint(ctx)
	}
}
