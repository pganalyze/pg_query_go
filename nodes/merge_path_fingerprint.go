// Auto-generated - DO NOT EDIT

package pg_query

import "io"

func (node MergePath) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "MergePath")

	for _, subNode := range node.Innersortkeys {
		subNode.Fingerprint(ctx)
	}

	for _, subNode := range node.Outersortkeys {
		subNode.Fingerprint(ctx)
	}

	for _, subNode := range node.PathMergeclauses {
		subNode.Fingerprint(ctx)
	}
}
