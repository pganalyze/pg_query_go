// Auto-generated - DO NOT EDIT

package pg_query

import "io"

func (node WithClause) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "WithClause")

	for _, subNode := range node.Ctes {
		subNode.Fingerprint(ctx)
	}
}
