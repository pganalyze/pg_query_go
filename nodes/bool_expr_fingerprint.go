// Auto-generated - DO NOT EDIT

package pg_query

import "io"

func (node BoolExpr) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "BoolExpr")

	for _, subNode := range node.Args {
		subNode.Fingerprint(ctx)
	}
}
