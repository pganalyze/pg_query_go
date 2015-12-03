// Auto-generated - DO NOT EDIT

package pg_query

import "io"

func (node CaseWhen) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "CaseWhen")
	if node.Expr != nil {
		node.Expr.Fingerprint(ctx)
	}

	if node.Result != nil {
		node.Result.Fingerprint(ctx)
	}
}
