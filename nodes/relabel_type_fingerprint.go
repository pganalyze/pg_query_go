// Auto-generated - DO NOT EDIT

package pg_query

import "io"

func (node RelabelType) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "RelabelType")
	if node.Arg != nil {
		node.Arg.Fingerprint(ctx)
	}
}
