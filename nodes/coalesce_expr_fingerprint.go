// Auto-generated - DO NOT EDIT

package pg_query

import "io"

func (node CoalesceExpr) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "COALESCE")

	for _, subNode := range node.Args {
		subNode.Fingerprint(ctx)
	}

	// Intentionally ignoring node.Location for fingerprinting
}
