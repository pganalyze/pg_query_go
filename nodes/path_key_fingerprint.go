// Auto-generated - DO NOT EDIT

package pg_query

import "io"

func (node PathKey) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "PathKey")
	if node.PkEclass != nil {
		node.PkEclass.Fingerprint(ctx)
	}
}
