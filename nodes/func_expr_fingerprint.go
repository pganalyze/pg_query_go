// Auto-generated - DO NOT EDIT

package pg_query

import "io"

func (node FuncExpr) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "FuncExpr")

	for _, subNode := range node.Args {
		subNode.Fingerprint(ctx)
	}
}
