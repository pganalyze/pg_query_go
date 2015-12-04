// Auto-generated - DO NOT EDIT

package pg_query

import "io"

func (node AlterFdwStmt) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "ALTERFDWSTMT")

	if node.Fdwname != nil {
		io.WriteString(ctx.hash, *node.Fdwname)
	}

	for _, subNode := range node.FuncOptions {
		subNode.Fingerprint(ctx)
	}

	for _, subNode := range node.Options {
		subNode.Fingerprint(ctx)
	}
}
