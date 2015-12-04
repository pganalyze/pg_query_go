// Auto-generated - DO NOT EDIT

package pg_query

import (
	"io"
	"strconv"
)

func (node NullTest) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "NULLTEST")

	if node.Arg != nil {
		node.Arg.Fingerprint(ctx)
	}

	io.WriteString(ctx.hash, strconv.FormatBool(node.Argisrow))
	io.WriteString(ctx.hash, strconv.Itoa(int(node.Nulltesttype)))
}
