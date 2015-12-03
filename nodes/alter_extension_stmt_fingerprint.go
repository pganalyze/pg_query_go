// Auto-generated - DO NOT EDIT

package pg_query

import "io"

func (node AlterExtensionStmt) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "AlterExtensionStmt")

	for _, subNode := range node.Options {
		subNode.Fingerprint(ctx)
	}
}
