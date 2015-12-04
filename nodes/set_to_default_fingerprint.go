// Auto-generated - DO NOT EDIT

package pg_query

import "io"

func (node SetToDefault) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "SETTODEFAULT")
	// Intentionally ignoring node.Location for fingerprinting
}
