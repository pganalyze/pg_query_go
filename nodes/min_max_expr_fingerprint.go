// Auto-generated - DO NOT EDIT

package pg_query

import (
	"io"
	"strconv"
)

func (node MinMaxExpr) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "MINMAX")

	for _, subNode := range node.Args {
		subNode.Fingerprint(ctx)
	}

	// Intentionally ignoring node.Location for fingerprinting

	io.WriteString(ctx.hash, strconv.Itoa(int(node.Op)))
}
