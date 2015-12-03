// Auto-generated - DO NOT EDIT

package pg_query

import "io"

func (node CreateFdwStmt) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "CreateFdwStmt")

	for _, subNode := range node.FuncOptions {
		subNode.Fingerprint(ctx)
	}

	for _, subNode := range node.Options {
		subNode.Fingerprint(ctx)
	}
}
