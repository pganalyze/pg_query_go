// Auto-generated - DO NOT EDIT

package pg_query

import "io"

func (node CoerceToDomainValue) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "COERCETODOMAINVALUE")
	// Intentionally ignoring node.Location for fingerprinting
}
