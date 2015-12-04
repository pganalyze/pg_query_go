// Auto-generated - DO NOT EDIT

package pg_query

import (
	"io"
	"strconv"
)

func (node MergePath) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "MERGEPATH")

	for _, subNode := range node.Innersortkeys {
		subNode.Fingerprint(ctx)
	}

	io.WriteString(ctx.hash, strconv.FormatBool(node.MaterializeInner))

	for _, subNode := range node.Outersortkeys {
		subNode.Fingerprint(ctx)
	}

	for _, subNode := range node.PathMergeclauses {
		subNode.Fingerprint(ctx)
	}
}
