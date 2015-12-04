// Auto-generated - DO NOT EDIT

package pg_query

import (
	"io"
	"strconv"
)

func (node Const) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "CONST")
	io.WriteString(ctx.hash, strconv.FormatBool(node.Constbyval))
	io.WriteString(ctx.hash, strconv.FormatBool(node.Constisnull))
	// Intentionally ignoring node.Location for fingerprinting
}
