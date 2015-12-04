// Auto-generated - DO NOT EDIT

package pg_query

import (
	"io"
	"strconv"
)

func (node CoerceViaIO) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "COERCEVIAIO")

	if node.Arg != nil {
		node.Arg.Fingerprint(ctx)
	}

	io.WriteString(ctx.hash, strconv.Itoa(int(node.Coerceformat)))
	// Intentionally ignoring node.Location for fingerprinting
}
