// Auto-generated - DO NOT EDIT

package pg_query

import "io"

func (node AlterRoleStmt) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "AlterRoleStmt")

	for _, subNode := range node.Options {
		subNode.Fingerprint(ctx)
	}
}
