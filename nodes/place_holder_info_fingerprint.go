// Auto-generated - DO NOT EDIT

package pg_query

import "io"

func (node PlaceHolderInfo) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "PLACEHOLDERINFO")

	if node.PhVar != nil {
		node.PhVar.Fingerprint(ctx)
	}
}
