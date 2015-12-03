// Auto-generated - DO NOT EDIT

package pg_query

import "io"

func (node TidScan) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "TidScan")

	for _, subNode := range node.Tidquals {
		subNode.Fingerprint(ctx)
	}
}
