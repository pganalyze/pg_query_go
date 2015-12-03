// Auto-generated - DO NOT EDIT

package pg_query

import "io"

func (node CreateForeignServerStmt) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "CreateForeignServerStmt")

	for _, subNode := range node.Options {
		subNode.Fingerprint(ctx)
	}
}
