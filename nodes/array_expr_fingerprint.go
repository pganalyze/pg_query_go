// Auto-generated - DO NOT EDIT

package pg_query

import (
	"io"
	"strconv"
)

func (node ArrayExpr) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "ARRAY")

	for _, subNode := range node.Elements {
		subNode.Fingerprint(ctx)
	}

	// Intentionally ignoring node.Location for fingerprinting

	io.WriteString(ctx.hash, strconv.FormatBool(node.Multidims))
}
