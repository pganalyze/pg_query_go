// Auto-generated - DO NOT EDIT

package pg_query

import "io"

func (node AccessPriv) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "AccessPriv")

	for _, subNode := range node.Cols {
		subNode.Fingerprint(ctx)
	}
}
