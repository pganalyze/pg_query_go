// Auto-generated - DO NOT EDIT

package pg_query

import "io"

func (node CreateRoleStmt) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "CreateRoleStmt")

	for _, subNode := range node.Options {
		subNode.Fingerprint(ctx)
	}
}
