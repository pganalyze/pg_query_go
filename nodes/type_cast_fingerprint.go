// Auto-generated - DO NOT EDIT

package pg_query

import "io"

func (node TypeCast) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "TYPECAST")

	if node.Arg != nil {
		node.Arg.Fingerprint(ctx)
	}

	// Intentionally ignoring node.Location for fingerprinting

	if node.TypeName != nil {
		node.TypeName.Fingerprint(ctx)
	}
}
