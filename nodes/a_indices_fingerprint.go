// Auto-generated - DO NOT EDIT

package pg_query

import "io"

func (node A_Indices) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "A_Indices")
	if node.Lidx != nil {
		node.Lidx.Fingerprint(ctx)
	}

	if node.Uidx != nil {
		node.Uidx.Fingerprint(ctx)
	}
}
