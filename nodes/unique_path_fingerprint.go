// Auto-generated - DO NOT EDIT

package pg_query

import "io"

func (node UniquePath) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "UniquePath")

	for _, subNode := range node.InOperators {
		subNode.Fingerprint(ctx)
	}

	if node.Subpath != nil {
		node.Subpath.Fingerprint(ctx)
	}

	for _, subNode := range node.UniqExprs {
		subNode.Fingerprint(ctx)
	}
}
