// Auto-generated - DO NOT EDIT

package pg_query

import "io"

func (node CreateExtensionStmt) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "CreateExtensionStmt")

	for _, subNode := range node.Options {
		subNode.Fingerprint(ctx)
	}
}
