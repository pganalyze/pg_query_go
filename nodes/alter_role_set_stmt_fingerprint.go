// Auto-generated - DO NOT EDIT

package pg_query

import "io"

func (node AlterRoleSetStmt) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "ALTERROLESETSTMT")

	if node.Database != nil {
		io.WriteString(ctx.hash, *node.Database)
	}

	if node.Role != nil {
		io.WriteString(ctx.hash, *node.Role)
	}

	if node.Setstmt != nil {
		node.Setstmt.Fingerprint(ctx)
	}
}
