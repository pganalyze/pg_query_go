// Auto-generated - DO NOT EDIT

package pg_query

import "io"

func (node MaterialPath) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "MaterialPath")
	if node.Subpath != nil {
		node.Subpath.Fingerprint(ctx)
	}
}
