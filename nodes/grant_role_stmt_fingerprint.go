// Auto-generated - DO NOT EDIT

package pg_query

import "io"

func (node GrantRoleStmt) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "GrantRoleStmt")

	for _, subNode := range node.GrantedRoles {
		subNode.Fingerprint(ctx)
	}

	for _, subNode := range node.GranteeRoles {
		subNode.Fingerprint(ctx)
	}
}
