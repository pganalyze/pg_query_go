// Auto-generated - DO NOT EDIT

package pg_query

import "io"

func (node AlterExtensionStmt) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "ALTEREXTENSIONSTMT")

	if node.Extname != nil {
		io.WriteString(ctx.hash, *node.Extname)
	}

	for _, subNode := range node.Options {
		subNode.Fingerprint(ctx)
	}
}
