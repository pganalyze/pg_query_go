// Auto-generated - DO NOT EDIT

package pg_query

import "io"

func (node HashPath) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "HashPath")

	for _, subNode := range node.PathHashclauses {
		subNode.Fingerprint(ctx)
	}
}
