// Auto-generated - DO NOT EDIT

package pg_query

import "io"

func (node LockStmt) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "LockStmt")

	for _, subNode := range node.Relations {
		subNode.Fingerprint(ctx)
	}
}
