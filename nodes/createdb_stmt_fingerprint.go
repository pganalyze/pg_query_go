// Auto-generated - DO NOT EDIT

package pg_query

import "io"

func (node CreatedbStmt) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "CreatedbStmt")

	for _, subNode := range node.Options {
		subNode.Fingerprint(ctx)
	}
}
