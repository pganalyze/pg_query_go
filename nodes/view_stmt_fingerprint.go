// Auto-generated - DO NOT EDIT

package pg_query

import "io"

func (node ViewStmt) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "ViewStmt")

	for _, subNode := range node.Aliases {
		subNode.Fingerprint(ctx)
	}

	for _, subNode := range node.Options {
		subNode.Fingerprint(ctx)
	}

	if node.Query != nil {
		node.Query.Fingerprint(ctx)
	}

	if node.View != nil {
		node.View.Fingerprint(ctx)
	}
}
