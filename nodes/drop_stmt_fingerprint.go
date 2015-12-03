// Auto-generated - DO NOT EDIT

package pg_query

import "io"

func (node DropStmt) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "DropStmt")

	for _, subNode := range node.Arguments {
		subNode.Fingerprint(ctx)
	}

	for _, subNode := range node.Objects {
		subNode.Fingerprint(ctx)
	}
}
