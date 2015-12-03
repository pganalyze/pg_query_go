// Auto-generated - DO NOT EDIT

package pg_query

import "io"

func (node A_Expr) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "A_Expr")
	if node.Lexpr != nil {
		node.Lexpr.Fingerprint(ctx)
	}

	for _, subNode := range node.Name {
		subNode.Fingerprint(ctx)
	}

	if node.Rexpr != nil {
		node.Rexpr.Fingerprint(ctx)
	}
}
