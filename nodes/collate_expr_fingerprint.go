// Auto-generated - DO NOT EDIT

package pg_query

import "io"

func (node CollateExpr) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "COLLATE")

	if node.Arg != nil {
		node.Arg.Fingerprint(ctx)
	}

	// Intentionally ignoring node.Location for fingerprinting
}
