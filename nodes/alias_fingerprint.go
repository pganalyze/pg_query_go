// Auto-generated - DO NOT EDIT

package pg_query

import "io"

func (node Alias) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "Alias")

	for _, subNode := range node.Colnames {
		subNode.Fingerprint(ctx)
	}
}
