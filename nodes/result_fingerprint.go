// Auto-generated - DO NOT EDIT

package pg_query

import "io"

func (node Result) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "RESULT")

	if node.Resconstantqual != nil {
		node.Resconstantqual.Fingerprint(ctx)
	}
}
