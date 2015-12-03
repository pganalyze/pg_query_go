// Auto-generated - DO NOT EDIT

package pg_query

import "io"

func (node AlterDatabaseStmt) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "AlterDatabaseStmt")

	for _, subNode := range node.Options {
		subNode.Fingerprint(ctx)
	}
}
