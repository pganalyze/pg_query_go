// Auto-generated - DO NOT EDIT

package pg_query

import "io"

func (node TidPath) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "TIDPATH")

	for _, subNode := range node.Tidquals {
		subNode.Fingerprint(ctx)
	}
}
