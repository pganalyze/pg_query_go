// Auto-generated - DO NOT EDIT

package pg_query

import "io"

func (node TypeCast) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "TypeCast")
	if node.Arg != nil {
		node.Arg.Fingerprint(ctx)
	}

	if node.TypeName != nil {
		node.TypeName.Fingerprint(ctx)
	}
}
