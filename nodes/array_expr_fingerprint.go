// Auto-generated - DO NOT EDIT

package pg_query

import "io"

func (node ArrayExpr) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "ArrayExpr")

	for _, subNode := range node.Elements {
		subNode.Fingerprint(ctx)
	}
}
