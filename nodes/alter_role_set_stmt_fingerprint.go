// Auto-generated - DO NOT EDIT

package pg_query

import "io"

func (node AlterRoleSetStmt) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "AlterRoleSetStmt")
	if node.Setstmt != nil {
		node.Setstmt.Fingerprint(ctx)
	}
}
