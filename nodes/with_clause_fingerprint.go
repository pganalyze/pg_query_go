// Auto-generated - DO NOT EDIT

package pg_query

import (
	"io"
	"strconv"
)

func (node WithClause) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "WITHCLAUSE")

	for _, subNode := range node.Ctes {
		subNode.Fingerprint(ctx)
	}

	// Intentionally ignoring node.Location for fingerprinting

	io.WriteString(ctx.hash, strconv.FormatBool(node.Recursive))
}
