// Auto-generated - DO NOT EDIT

package pg_query

import (
	"io"
	"strconv"
)

func (node RowExpr) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "ROW")

	for _, subNode := range node.Args {
		subNode.Fingerprint(ctx)
	}

	for _, subNode := range node.Colnames {
		subNode.Fingerprint(ctx)
	}

	// Intentionally ignoring node.Location for fingerprinting

	io.WriteString(ctx.hash, strconv.Itoa(int(node.RowFormat)))
}
