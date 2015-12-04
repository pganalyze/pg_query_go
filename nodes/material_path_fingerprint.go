// Auto-generated - DO NOT EDIT

package pg_query

import "io"

func (node MaterialPath) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "MATERIALPATH")

	if node.Subpath != nil {
		node.Subpath.Fingerprint(ctx)
	}
}
