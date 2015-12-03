// Auto-generated - DO NOT EDIT

package pg_query

import "io"

func (node CoalesceExpr) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "CoalesceExpr")

	for _, subNode := range node.Args {
		subNode.Fingerprint(ctx)
	}
}
