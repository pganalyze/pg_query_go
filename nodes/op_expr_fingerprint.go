// Auto-generated - DO NOT EDIT

package pg_query

import "io"

func (node OpExpr) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "OpExpr")

	for _, subNode := range node.Args {
		subNode.Fingerprint(ctx)
	}
}
