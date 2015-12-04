// Auto-generated - DO NOT EDIT

package pg_query

import "io"

func (node AlterRoleStmt) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "ALTERROLESTMT")

	for _, subNode := range node.Options {
		subNode.Fingerprint(ctx)
	}

	if node.Role != nil {
		io.WriteString(ctx.hash, *node.Role)
	}
}
