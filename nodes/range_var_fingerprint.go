// Auto-generated - DO NOT EDIT

package pg_query

import "io"

func (node RangeVar) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "RangeVar")
	if node.Alias != nil {
		node.Alias.Fingerprint(ctx)
	}
}
