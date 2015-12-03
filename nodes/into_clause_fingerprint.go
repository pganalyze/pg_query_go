// Auto-generated - DO NOT EDIT

package pg_query

import "io"

func (node IntoClause) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "IntoClause")

	for _, subNode := range node.ColNames {
		subNode.Fingerprint(ctx)
	}

	for _, subNode := range node.Options {
		subNode.Fingerprint(ctx)
	}

	if node.Rel != nil {
		node.Rel.Fingerprint(ctx)
	}

	if node.ViewQuery != nil {
		node.ViewQuery.Fingerprint(ctx)
	}
}
