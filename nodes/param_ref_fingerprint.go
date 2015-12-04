// Auto-generated - DO NOT EDIT

package pg_query

import "io"

func (node ParamRef) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "PARAMREF")
	// Intentionally ignoring node.Location for fingerprinting
}
