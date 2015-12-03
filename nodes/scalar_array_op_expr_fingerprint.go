// Auto-generated - DO NOT EDIT

package pg_query

import "io"

func (node ScalarArrayOpExpr) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "ScalarArrayOpExpr")

	for _, subNode := range node.Args {
		subNode.Fingerprint(ctx)
	}
}
