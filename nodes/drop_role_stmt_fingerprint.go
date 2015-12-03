// Auto-generated - DO NOT EDIT

package pg_query

import "io"

func (node DropRoleStmt) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "DropRoleStmt")

	for _, subNode := range node.Roles {
		subNode.Fingerprint(ctx)
	}
}
