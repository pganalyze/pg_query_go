// Auto-generated - DO NOT EDIT

package pg_query

import "io"

func (node HashJoin) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "HASHJOIN")

	for _, subNode := range node.Hashclauses {
		subNode.Fingerprint(ctx)
	}
}
