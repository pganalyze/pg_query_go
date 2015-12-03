// Auto-generated - DO NOT EDIT

package pg_query

import "io"

func (node ExecuteStmt) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "ExecuteStmt")

	for _, subNode := range node.Params {
		subNode.Fingerprint(ctx)
	}
}
