// Auto-generated - DO NOT EDIT

package pg_query

import "io"

func (node CoerceViaIO) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "CoerceViaIO")
	if node.Arg != nil {
		node.Arg.Fingerprint(ctx)
	}
}
