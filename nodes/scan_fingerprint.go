// Auto-generated - DO NOT EDIT

package pg_query

import "io"

func (node Scan) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "SCAN")
}
