// Auto-generated - DO NOT EDIT

package pg_query

import "io"

func (node FieldSelect) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "FieldSelect")
	if node.Arg != nil {
		node.Arg.Fingerprint(ctx)
	}
}
