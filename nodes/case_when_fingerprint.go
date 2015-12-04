// Auto-generated - DO NOT EDIT

package pg_query

import "io"

func (node CaseWhen) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "WHEN")

	if node.Expr != nil {
		node.Expr.Fingerprint(ctx)
	}

	// Intentionally ignoring node.Location for fingerprinting

	if node.Result != nil {
		node.Result.Fingerprint(ctx)
	}
}
