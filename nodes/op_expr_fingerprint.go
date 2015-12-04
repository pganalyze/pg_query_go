// Auto-generated - DO NOT EDIT

package pg_query

import (
	"io"
	"strconv"
)

func (node OpExpr) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "OPEXPR")

	for _, subNode := range node.Args {
		subNode.Fingerprint(ctx)
	}

	// Intentionally ignoring node.Location for fingerprinting

	io.WriteString(ctx.hash, strconv.FormatBool(node.Opretset))
}
