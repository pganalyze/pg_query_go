// Auto-generated - DO NOT EDIT

package pg_query

import (
	"io"
	"strconv"
)

func (node ConstraintsSetStmt) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "CONSTRAINTSSETSTMT")

	for _, subNode := range node.Constraints {
		subNode.Fingerprint(ctx)
	}

	io.WriteString(ctx.hash, strconv.FormatBool(node.Deferred))
}
