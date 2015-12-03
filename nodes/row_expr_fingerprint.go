// Auto-generated - DO NOT EDIT

package pg_query

import "io"

func (node RowExpr) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "RowExpr")

	for _, subNode := range node.Args {
		subNode.Fingerprint(ctx)
	}

	for _, subNode := range node.Colnames {
		subNode.Fingerprint(ctx)
	}
}
