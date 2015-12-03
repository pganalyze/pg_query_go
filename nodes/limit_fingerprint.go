// Auto-generated - DO NOT EDIT

package pg_query

import "io"

func (node Limit) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "Limit")
	if node.LimitCount != nil {
		node.LimitCount.Fingerprint(ctx)
	}

	if node.LimitOffset != nil {
		node.LimitOffset.Fingerprint(ctx)
	}
}
