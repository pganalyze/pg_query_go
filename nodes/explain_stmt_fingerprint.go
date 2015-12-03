// Auto-generated - DO NOT EDIT

package pg_query

import "io"

func (node ExplainStmt) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "ExplainStmt")

	for _, subNode := range node.Options {
		subNode.Fingerprint(ctx)
	}

	if node.Query != nil {
		node.Query.Fingerprint(ctx)
	}
}
