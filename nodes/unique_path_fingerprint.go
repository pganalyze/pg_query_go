// Auto-generated - DO NOT EDIT

package pg_query

import (
	"io"
	"strconv"
)

func (node UniquePath) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "UNIQUEPATH")

	for _, subNode := range node.InOperators {
		subNode.Fingerprint(ctx)
	}

	if node.Subpath != nil {
		node.Subpath.Fingerprint(ctx)
	}

	io.WriteString(ctx.hash, strconv.Itoa(int(node.Umethod)))

	for _, subNode := range node.UniqExprs {
		subNode.Fingerprint(ctx)
	}
}
