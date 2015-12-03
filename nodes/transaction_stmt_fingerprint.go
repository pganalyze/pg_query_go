// Auto-generated - DO NOT EDIT

package pg_query

import "io"

func (node TransactionStmt) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "TransactionStmt")

	for _, subNode := range node.Options {
		subNode.Fingerprint(ctx)
	}
}
