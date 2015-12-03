// Auto-generated - DO NOT EDIT

package pg_query

import "io"

func (node Join) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "Join")

	for _, subNode := range node.Joinqual {
		subNode.Fingerprint(ctx)
	}
}
