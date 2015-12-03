// Auto-generated - DO NOT EDIT

package pg_query

import "io"

func (node FromExpr) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "FromExpr")

	for _, subNode := range node.Fromlist {
		subNode.Fingerprint(ctx)
	}

	if node.Quals != nil {
		node.Quals.Fingerprint(ctx)
	}
}
