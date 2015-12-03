// Auto-generated - DO NOT EDIT

package pg_query

import "io"

func (node DefElem) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "DefElem")
	if node.Arg != nil {
		node.Arg.Fingerprint(ctx)
	}
}
