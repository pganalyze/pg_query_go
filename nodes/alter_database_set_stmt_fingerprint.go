// Auto-generated - DO NOT EDIT

package pg_query

import "io"

func (node AlterDatabaseSetStmt) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "AlterDatabaseSetStmt")
	if node.Setstmt != nil {
		node.Setstmt.Fingerprint(ctx)
	}
}
