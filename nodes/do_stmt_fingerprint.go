// Auto-generated - DO NOT EDIT

package pg_query

import "io"

func (node DoStmt) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "DoStmt")

	for _, subNode := range node.Args {
		subNode.Fingerprint(ctx)
	}
}
