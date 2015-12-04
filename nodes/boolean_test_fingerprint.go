// Auto-generated - DO NOT EDIT

package pg_query

import (
	"io"
	"strconv"
)

func (node BooleanTest) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "BOOLEANTEST")

	if node.Arg != nil {
		node.Arg.Fingerprint(ctx)
	}

	io.WriteString(ctx.hash, strconv.Itoa(int(node.Booltesttype)))
}
