// Auto-generated - DO NOT EDIT

package pg_query

import (
	"io"
	"strconv"
)

func (node ConvertRowtypeExpr) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "CONVERTROWTYPEEXPR")

	if node.Arg != nil {
		node.Arg.Fingerprint(ctx)
	}

	io.WriteString(ctx.hash, strconv.Itoa(int(node.Convertformat)))
	// Intentionally ignoring node.Location for fingerprinting
}
