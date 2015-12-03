// Auto-generated - DO NOT EDIT

package pg_query

import "io"

func (node AlterForeignServerStmt) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "AlterForeignServerStmt")

	for _, subNode := range node.Options {
		subNode.Fingerprint(ctx)
	}
}
