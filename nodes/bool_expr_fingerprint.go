// Auto-generated - DO NOT EDIT

package pg_query

import (
	"io"
	"strconv"
)

func (node BoolExpr) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "BOOLEXPR")

	for _, subNode := range node.Args {
		subNode.Fingerprint(ctx)
	}

	io.WriteString(ctx.hash, strconv.Itoa(int(node.Boolop)))
	// Intentionally ignoring node.Location for fingerprinting
}
