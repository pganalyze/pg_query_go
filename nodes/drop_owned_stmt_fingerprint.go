// Auto-generated - DO NOT EDIT

package pg_query

import "io"

func (node DropOwnedStmt) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "DropOwnedStmt")

	for _, subNode := range node.Roles {
		subNode.Fingerprint(ctx)
	}
}
