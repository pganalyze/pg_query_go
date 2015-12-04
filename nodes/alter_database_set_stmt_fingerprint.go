// Auto-generated - DO NOT EDIT

package pg_query

import "io"

func (node AlterDatabaseSetStmt) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "ALTERDATABASESETSTMT")

	if node.Dbname != nil {
		io.WriteString(ctx.hash, *node.Dbname)
	}

	if node.Setstmt != nil {
		node.Setstmt.Fingerprint(ctx)
	}
}
