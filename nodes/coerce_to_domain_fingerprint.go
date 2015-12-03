// Auto-generated - DO NOT EDIT

package pg_query

import "io"

func (node CoerceToDomain) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "CoerceToDomain")
	if node.Arg != nil {
		node.Arg.Fingerprint(ctx)
	}
}
