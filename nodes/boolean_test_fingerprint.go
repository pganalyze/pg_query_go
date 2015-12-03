// Auto-generated - DO NOT EDIT

package pg_query

import "io"

func (node BooleanTest) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "BooleanTest")
	if node.Arg != nil {
		node.Arg.Fingerprint(ctx)
	}
}
