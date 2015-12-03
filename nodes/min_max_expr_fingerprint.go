// Auto-generated - DO NOT EDIT

package pg_query

import "io"

func (node MinMaxExpr) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "MinMaxExpr")

	for _, subNode := range node.Args {
		subNode.Fingerprint(ctx)
	}
}
