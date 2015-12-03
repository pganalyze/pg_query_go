// Auto-generated - DO NOT EDIT

package pg_query

import "io"

func (node WithCheckOption) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "WithCheckOption")
	if node.Qual != nil {
		node.Qual.Fingerprint(ctx)
	}
}
