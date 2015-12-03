// Auto-generated - DO NOT EDIT

package pg_query

import "io"

func (node RangeSubselect) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "RangeSubselect")
	if node.Alias != nil {
		node.Alias.Fingerprint(ctx)
	}

	if node.Subquery != nil {
		node.Subquery.Fingerprint(ctx)
	}
}
