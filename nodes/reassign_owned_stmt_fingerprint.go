// Auto-generated - DO NOT EDIT

package pg_query

import "io"

func (node ReassignOwnedStmt) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "ReassignOwnedStmt")

	for _, subNode := range node.Roles {
		subNode.Fingerprint(ctx)
	}
}
