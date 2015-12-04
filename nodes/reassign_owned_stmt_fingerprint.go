// Auto-generated - DO NOT EDIT

package pg_query

import "io"

func (node ReassignOwnedStmt) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "REASSIGNOWNEDSTMT")

	if node.Newrole != nil {
		io.WriteString(ctx.hash, *node.Newrole)
	}

	for _, subNode := range node.Roles {
		subNode.Fingerprint(ctx)
	}
}
