// Auto-generated - DO NOT EDIT

package pg_query

import "io"

func (node NullTest) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "NullTest")
	if node.Arg != nil {
		node.Arg.Fingerprint(ctx)
	}
}
