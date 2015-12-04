// Auto-generated - DO NOT EDIT

package pg_query

import (
	"io"
	"strconv"
)

func (node Join) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "JOIN")

	for _, subNode := range node.Joinqual {
		subNode.Fingerprint(ctx)
	}

	io.WriteString(ctx.hash, strconv.Itoa(int(node.Jointype)))
}
