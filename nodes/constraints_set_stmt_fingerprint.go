// Auto-generated - DO NOT EDIT

package pg_query

import "io"

func (node ConstraintsSetStmt) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "ConstraintsSetStmt")

	for _, subNode := range node.Constraints {
		subNode.Fingerprint(ctx)
	}
}
