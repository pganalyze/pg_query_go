// Auto-generated - DO NOT EDIT

package pg_query

import (
	"io"
	"strconv"
)

func (node ArrayCoerceExpr) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "ARRAYCOERCEEXPR")

	if node.Arg != nil {
		node.Arg.Fingerprint(ctx)
	}

	io.WriteString(ctx.hash, strconv.Itoa(int(node.Coerceformat)))
	io.WriteString(ctx.hash, strconv.FormatBool(node.IsExplicit))
	// Intentionally ignoring node.Location for fingerprinting
}
