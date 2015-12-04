// Auto-generated - DO NOT EDIT

package pg_query

import "io"

func (node NamedArgExpr) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "NAMEDARGEXPR")

	if node.Arg != nil {
		node.Arg.Fingerprint(ctx)
	}

	// Intentionally ignoring node.Location for fingerprinting

	if node.Name != nil {
		io.WriteString(ctx.hash, *node.Name)
	}
}
